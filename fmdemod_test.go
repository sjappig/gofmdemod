package fmdemod

import (
    "testing"
    "math"
    "math/cmplx"
)

func TestFmDemodulation(t *testing.T) {
    f := float64(3)
    fs := float64(10)
    prevSample := cmplx.Exp(0.0)
    for idx := 1; idx < 10; idx ++ {
        currSample := cmplx.Exp(complex(0, 2.0 * math.Pi * f * float64(idx) / fs))
        actual := fmDemodulation(currSample, prevSample)
        expected := f/fs
        if actual != expected {
            t.Fatalf(`Got %f, expected %f for index %d`, actual, expected, idx)
        }
        prevSample = currSample
    }
}

func TestToComplex128(t *testing.T) {
    actual := toComplex128(0, 255)
    expected := -128 + 127i
    if actual != expected {
        t.Fatalf(`Got %f, expected %f`, actual, expected)
    }
}

func TestQuantizeToUin16(t *testing.T) {
    actual := quantizeToUint16(0.9999999)
    expected := uint16(65535)
    if actual != expected {
        t.Fatalf(`Got %d, expected %d`, actual, expected)
    }
}

func TestToLittleEndianBytes(t *testing.T) {
    actualFirst, actualSecond := toLittleEndianBytes(65535)
    expectedFirst, expectedSecond := byte(255), byte(255)
    if actualFirst != expectedFirst {
        t.Fatalf(`First byte: Got %d, expected %d`, actualFirst, expectedFirst)
    }
    if actualSecond != expectedSecond {
        t.Fatalf(`Second byte: Got %d, expected %d`, actualSecond, expectedSecond)
    }
}
