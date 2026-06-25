package draw

import (
	"image/color"
	"testing"
)

var resultInterface color.Color
var resultStruct color.RGBA

func BenchmarkOver_Interface(b *testing.B) {
	baseC := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	var a1, r1, g1, b1 uint32 = 128, 50, 50, 50

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Assign to interface (will allocate)
		resultInterface = Over(&baseC, a1, r1, g1, b1)
	}
}

func BenchmarkOver_Struct(b *testing.B) {
	baseC := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	var a1, r1, g1, b1 uint32 = 128, 50, 50, 50

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Assign to struct (should NOT allocate)
		resultStruct = Over(&baseC, a1, r1, g1, b1)
	}
}

func TestOverMath(t *testing.T) {
	// Base is opaque grey: r=100, g=100, b=100, a=255. In 16-bit space, 100 * 257 = 25700. a=65535.
	baseColor := color.RGBA{100, 100, 100, 255}

	// Overlay is opaque Amiga blue: r=0, g=80, b=160, a=255. In 16-bit space: g=20560, b=41120, a=65535.
	overlayColor := color.RGBA{0, 80, 160, 255}

	// The overlay values from RGBA() are pre-multiplied by alpha.
	r, g, b, a := overlayColor.RGBA()

	res := Over(baseColor, a, r, g, b)

	if res.R != 0 || res.G != 80 || res.B != 160 || res.A != 255 {
		t.Errorf("Expected {0, 80, 160, 255}, got %v. Math overflow likely occurred.", res)
	}
}
