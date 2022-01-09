# FM demodulation in Go

Requires rtlsdr.

Server:

    rtl_sdr -f 91900000 -s 256000 - | go run fmdemod.go | nc -l 1337


Client:

    nc localhost 1337 | play -t raw -r 256k -e unsigned -b 16 -c 1 -V1 -L -
