# dubstep-data

The programs I wrote for the `Encoding data in dubstep drops` blog post: https://blog.benjojo.co.uk/post/encoding-data-into-dubstep-drops

**Note this does not work on Windows machines at this point in time**

## Packages Required ## 
* [SoX (Sound Exchange)](http://sox.sourceforge.net/) -  Convert various formats of computer audio files to other formats. Also applies various effects to these sound files
    * `sudo apt-get install sox`
* [Google Golang](https://golang.org/) -  Google's Go Language
    * `sudo apt-get install golang`
    * Note: You will need to set the GOPATH
* [FFMPEG](https://ffmpeg.zeranoe.com/) - Decode, encode, filter and play files
    * `sudo apt-get install ffmpeg`

## How to run ##
* Note: If you have not set your GOPATH correctly, it will not be able to perform the `go get` commmands

Acquire repos:
``` 
go get github.com/dgryski/go-bitstream
go get github.com/benjojo/dubstep-data
```
Once they have downloaded, navigate to the 'github.com' folder. By default it can be found through the home/ directory

```
cd benjojo
cd dubstep-data
go env
go build
```

Once go has compiled the main file, you will need to edit the bash script 'prep-wav.sh'.

Change
`./ASK-dubstep -input in.f64.data -data "$2"`
to
`./dubstep-data -input in.f64.data -data "$2"`

This will be fixed at a later date.

``` ./prep-wav.sh "*.wav" "`date`" ```

If you encounter an error with Sox, saying that the "filter frequency must be less than sample-rate / 2", this might be an indication that the wav file is corrupt or too small.

The script file will complete, creating new files, most importantly the 'encoded-bassline.wav'.

Move to the decode folder
```
cd decode
go build
ffmpeg -i ../encoded-bassline.wav -f f64le -ar 44100 -ac 1 -y in.f64.data
```
To show steganograph:
```
mkfifo sound
ffplay -f f64le -ar 44100 -ac 1 -i sound & cat in.f64.data | tee sound | ./decode -input /dev/stdin
```
Otherwise run `cat in.f64.data | tee sound | ./decode -input /dev/stdin`

To view the decoded message:
`./decode`
```
