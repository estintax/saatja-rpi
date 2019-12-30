# saatja

A small program for the Raspberry Pi to transmit Morse code using the FS1000A RF-transmitter.

![FS1000A](https://raw.githubusercontent.com/estintax/saatja/master/doc/fs1000a.png)

## Usage

```Bash
# apt-get install golang
$ git clone https://github.com/estintax/saatja.git
$ cd saatja
$ git submodule init
$ git submodule update
$ go build
$ ./saatja [text (not case sensitive)]
```

## Warnings

1. The FS1000A transmitter broadcasts at approximately 433.986 MHz. This is a free frequency in many countries. However, I am not sure that this frequency is free in all countries. Perhaps, before using the program, you will need to obtain permission in advance to use this frequency.
2. The frequency of the FS1000A transmitter is not static. For me personally, the frequency changes a little, which makes the tone of the signal dynamic.
