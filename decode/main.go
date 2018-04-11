package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	bitstream "github.com/dgryski/go-bitstream"
)

func main() {
	ins := flag.String("input", "./in.f64.data", "")
	symbolrate := flag.Int("srate", 5500, "Symbol rate in samples")
	flag.Parse()

	inf, err := os.OpenFile(*ins, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Unable to open output file %s", err.Error())
	}

	bw := bitstream.NewWriter(os.Stdout)
	bw.WriteBit(bitstream.Bit(false))
	// fmt.Print("0")

	SamplesUntilChange := *symbolrate
	bits := 1
	negscore := 0

	for {
		var raws float64
		err := binary.Read(inf, binary.LittleEndian, &raws)
		if err != nil {
			fmt.Print("\n")
			log.Printf("Leaving %s", err.Error())
			break
		}

		isNeg := raws < 0
		if isNeg {
			negscore++
		}

		SamplesUntilChange--
		if SamplesUntilChange == 0 {
			rsp := negscore < (*symbolrate / 2)
			if bits != 1 {
				// log.Printf("NS: %d | %v", negscore, bool(rsp))
				bw.WriteBit(bitstream.Bit(rsp))
				if rsp {
					// fmt.Print("1")
				} else {
					// fmt.Print("0")
				}
			}

			negscore = 0
			SamplesUntilChange = *symbolrate
			bits++
		}

	}
	log.Printf("Finished with %d bits / %d bytes", bits, bits/8)
}
