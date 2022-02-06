package oklab

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestOklabConversion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	var minL, maxL, minA, maxA, minB, maxB float64
	for r := 0; r <= 0xff; r++ {
		for g := 0; g <= 0xff; g++ {
			for b := 0; b <= 0xff; b++ {
				start := color.NRGBA{uint8(r), uint8(g), uint8(b), 0xff}
				oklab := OklabModel.Convert(start).(Oklab)
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
				oklch := OklchModel.Convert(start).(Oklch)
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

func TestGradientImage(t *testing.T) {
	f, err := os.Create("gradient.png")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	if err := png.Encode(f, Gradient{}); err != nil {
		t.Error(err)
	}
}

type Gradient struct{}

func (g Gradient) ColorModel() color.Model {
	return OklabModel
}

func (g Gradient) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1200, 600)
}

func (g Gradient) At(x, y int) color.Color {
	return Oklab{0.8, lerpA(float64(x) / float64(g.Bounds().Dx())), lerpB(float64(y) / float64(g.Bounds().Dy()))}
}

func lerpA(x float64) float64 {
	minA, maxA := -0.233888, 0.276216
	return x*(maxA-minA) + minA
}

func lerpB(x float64) float64 {
	minB, maxB := -0.311528, 0.198570
	return x*(maxB-minB) + minB
}
