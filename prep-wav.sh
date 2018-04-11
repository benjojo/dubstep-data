#!/bin/bash

sox $1 prep-lower.wav sinc 0k-0.1k
sox $1 prep-higher.wav sinc 0.1k-22k
ffmpeg -i prep-lower.wav -f f64le -ar 44100 -ac 1 -y in.f64.data 2> /dev/null

./ASK-dubstep -input in.f64.data -data "$2"

ffmpeg -f f64le -ar 44100 -ac 1 -i out.f64.data -y encoded-bass.wav 2> /dev/null
sox -m encoded-bass.wav prep-higher.wav encoded-bassline.wav

echo "Output saved, listen with $ vlc encoded-bassline.wav"
