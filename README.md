# FM demodulation in Go

Reads complex uint8 (re, im) from stdin, outputs UQ16 (unsigned fixed point, 16 fraction bits).

## Receiving radio broadcast

Requires rtlsdr (software and a receiver dongle) and sox (for audio output). For oneliner without
server-client-architecture, drop the *nc* and just pipe the output of the demodulator to *play*.

Server:

    rtl_sdr -f 91900000 -s 256000 - | go run fmdemod.go | nc -l 1337


Client:

    nc localhost 1337 | play -t raw -r 256k -e unsigned -b 16 -c 1 -V1 -L -
