package oklab_test

import (
	"fmt"
	"github.com/alltom/oklab"
	"image/color"
	"testing"
)

func Example_convertToOklab() {
	rgbc := color.RGBA{0xff, 0xdf, 0xe7, 0xff}
	oklabc := oklab.OklabModel.Convert(rgbc).(oklab.Oklab)
	fmt.Printf("L: %.2f, a: %.2f, b: %.2f\n", oklabc.L, oklabc.A, oklabc.B)
	// Output: L: 0.93, a: 0.04, b: 0.00
}

func Example_convertOklabToRGB() {
	oklabc := oklab.Oklab{L: 0.9322421414586456, A: 0.03673270292094283, B: 0.0006123556644819055}
	r, g, b, _ := oklabc.RGBA()
	fmt.Printf("R: 0x%x, G: 0x%x, B: 0x%x\n", r>>8, g>>8, b>>8)
	// Output: R: 0xff, G: 0xdf, B: 0xe7
}

func TestOklabConversion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	var minL, maxL, minA, maxA, minB, maxB float64
	for r := 0; r <= 0xff; r++ {
		for g := 0; g <= 0xff; g++ {
			for b := 0; b <= 0xff; b++ {
				start := color.NRGBA{uint8(r), uint8(g), uint8(b), 0xff}
				oklab := oklab.OklabModel.Convert(start).(oklab.Oklab)
				final := color.NRGBAModel.Convert(oklab)
				if start != final {
					t.Errorf("rgb(oklab(%v)) = %v; want %v", start, final, start)
				}
				if r == 0 && g == 0 && b == 0 {
					minL, maxL = oklab.L, oklab.L
					minA, maxA = oklab.A, oklab.A
					minB, maxB = oklab.B, oklab.B
				}
				if oklab.L < minL {
					minL = oklab.L
				}
				if oklab.L > maxL {
					maxL = oklab.L
				}
				if oklab.A < minA {
					minA = oklab.A
				}
				if oklab.A > maxA {
					maxA = oklab.A
				}
				if oklab.B < minB {
					minB = oklab.B
				}
				if oklab.B > maxB {
					maxB = oklab.B
				}
			}
		}
	}
	t.Logf("L: %f–%f", minL, maxL)
	t.Logf("A: %f–%f", minA, maxA)
	t.Logf("B: %f–%f", minB, maxB)
}

func TestOklchConversion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	var minL, maxL, minC, maxC, minH, maxH float64
	for r := 0; r <= 0xff; r++ {
		for g := 0; g <= 0xff; g++ {
			for b := 0; b <= 0xff; b++ {
				start := color.NRGBA{uint8(r), uint8(g), uint8(b), 0xff}
				oklch := oklab.OklchModel.Convert(start).(oklab.Oklch)
				final := color.NRGBAModel.Convert(oklch)
				if start != final {
					t.Errorf("rgb(oklch(%v)) = %v; want %v", start, final, start)
				}
				if r == 0 && g == 0 && b == 0 {
					minL, maxL = oklch.L, oklch.L
					minC, maxC = oklch.C, oklch.C
					minH, maxH = oklch.H, oklch.H
				}
				if oklch.L < minL {
					minL = oklch.L
				}
				if oklch.L > maxL {
					maxL = oklch.L
				}
				if oklch.C < minC {
					minC = oklch.C
				}
				if oklch.C > maxC {
					maxC = oklch.C
				}
				if oklch.H < minH {
					minH = oklch.H
				}
				if oklch.H > maxH {
					maxH = oklch.H
				}
			}
		}
	}
	t.Logf("L: %f–%f", minL, maxL)
	t.Logf("C: %f–%f", minC, maxC)
	t.Logf("H: %f–%f", minH, maxH)
}
