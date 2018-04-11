package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dgryski/go-bitstream"
)

func main() {
	ins := flag.String("input", "./in.f64.data", "")
	outs := flag.String("out", "./out.f64.data", "")
	encodetarget := flag.String("data", "Hello World!", "")
	symbolrate := flag.Int("srate", 5500, "Symbol rate in samples")

	flag.Parse()

	inf, err := os.OpenFile(*ins, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Unable to open output file %s", err.Error())
	}

	outf, err := os.OpenFile(*outs, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Unable to open output file %s", err.Error())
	}

	sr := strings.NewReader(*encodetarget)
	bitreader := bitstream.NewReader(sr)

	SamplesUntilChange := *symbolrate
	firstbit, _ := bitreader.ReadBit()
	secondbit, _ := bitreader.ReadBit()
	UpperFlip := bool(firstbit)
	nextbit := bool(secondbit)
	bits := 0

	for {
		var raws float64
		err := binary.Read(inf, binary.LittleEndian, &raws)
		if err != nil {
			fmt.Print("\n")
			log.Printf("Leaving %s", err.Error())
			break
		}

		// First, obtain a upper flip
		normie := (raws + 1) / 2

		SamplesUntilChange--
		if SamplesUntilChange == 0 {
			UpperFlip = nextbit
			b, _ := bitreader.ReadBit()
			nextbit = bool(b)
			if UpperFlip {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}

			SamplesUntilChange = *symbolrate
			bits++
		}

		if !UpperFlip {
			normie = normie * -1
		}

		if SamplesUntilChange < 1000 && UpperFlip != nextbit {
			dest := normie * -1
			normie = lerp(dest, normie, float64(SamplesUntilChange)/1000.0)
		}

		binary.Write(outf, binary.LittleEndian, normie)
	}
	log.Printf("Finished with %d bits / %d bytes", bits, bits/8)
}

func lerp(a, b, n float64) float64 {
	return (1-n)*a + n*b
}
