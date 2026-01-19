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
