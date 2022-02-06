package oklab_test

import (
	"fmt"
	"github.com/alltom/oklab"
	"image"
	"image/color"
	"image/png"
	"os"
)

func Example_gradientImage() {
	f, _ := os.Create("gradient.png")
	png.Encode(f, GradientImage{})
	fmt.Println("err =", f.Close())
	// Output: err = <nil>
}

type GradientImage struct{}

func (g GradientImage) ColorModel() color.Model {
	return oklab.OklabModel
}

func (g GradientImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1200, 600)
}

func (g GradientImage) At(x, y int) color.Color {
	a := lerp(float64(x)/float64(g.Bounds().Dx()), -0.233888, 0.276216)
	b := lerp(float64(y)/float64(g.Bounds().Dy()), -0.311528, 0.198570)
	return oklab.Oklab{0.8, a, b}
}

func lerp(x, min, max float64) float64 {
	return x*(max-min) + min
}
