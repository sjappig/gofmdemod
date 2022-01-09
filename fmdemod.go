package main

import (
    "bufio"
    "log"
    "math"
    "math/cmplx"
    "os"
)

const SampleCount = 8192

func fmDemodulation(currSample, prevSample complex128) float64 {
    return cmplx.Phase(currSample * cmplx.Conj(prevSample)) / math.Pi
}

func toComplex128(re, im uint8) complex128 {
    return complex(float64(re) - 128, float64(im) - 128)
}

func quantizeToUint16(value float64) uint16 {
    value = math.Max(value, -1)
    value = math.Min(value, 1)
    return uint16((value + 1) * 32768)
}

func toLittleEndianBytes(value uint16) (byte, byte) {
    return byte(value & 0x00FF), byte((value & 0xFF00) >> 8)
}

func main() {
    iqByteBuffer := make([]byte, SampleCount * 2)
    writeBuffer := make([]byte, SampleCount * 2)

    prevSample := complex(0, 0)

    reader := bufio.NewReader(os.Stdin)

    for {
        readCount, err := reader.Read(iqByteBuffer)

        if err != nil {
            log.Fatal(err)
        }

        if readCount == 0 {
            break
        }

        for idx := 0; idx < readCount; idx += 2 {
            currSample := toComplex128(iqByteBuffer[idx], iqByteBuffer[idx + 1])
            writeBuffer[idx], writeBuffer[idx + 1] = toLittleEndianBytes(quantizeToUint16(fmDemodulation(currSample, prevSample)))
            prevSample = currSample
        }
        os.Stdout.Write(writeBuffer[:readCount])
    }
}
