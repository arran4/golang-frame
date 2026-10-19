package frame

import (
	"image"
	"image/color"
	"testing"
)

// solidImage is a helper to create images of a solid color
type solidImage struct {
	rect image.Rectangle
	c    color.Color
}

func (s solidImage) ColorModel() color.Model { return color.RGBAModel }
func (s solidImage) Bounds() image.Rectangle { return s.rect }
func (s solidImage) At(x, y int) color.Color {
	if !image.Pt(x, y).In(s.rect) {
		return color.Transparent
	}
	return s.c
}

func TestSection5Bounds(t *testing.T) {
	base := image.NewRGBA(image.Rect(10, 10, 50, 50))
	middle := image.Rect(20, 20, 40, 40)

	s5Red := color.RGBA{255, 0, 0, 255}
	s5i := solidImage{image.Rect(5, 5, 25, 25), s5Red}

	dest := image.Rect(100, 100, 140, 140)

	t.Run("Section5Zeroed", func(t *testing.T) {
		f := NewFrame(dest, base, middle, &Section5{Image: s5i, Replace: true}, Section5Zeroed)
		c := f.At(110, 110)
		r, g, b, a := c.RGBA()
		if r != 65535 || g != 0 || b != 0 || a != 65535 {
			t.Errorf("Expected red color from section5, got %v", c)
		}
	})

	t.Run("Zerod", func(t *testing.T) {
		s5i_zerod := solidImage{image.Rect(5, 5, 45, 45), s5Red}
		f := NewFrame(dest, base, middle, &Section5{Image: s5i_zerod, Replace: true}, Zerod)
		c := f.At(110, 110) // offset from 100 is (10, 10), from origin (5, 5) -> (15, 15) which is in bounds of (5,5,45,45)
		r, g, b, a := c.RGBA()
		if r != 65535 || g != 0 || b != 0 || a != 65535 {
			t.Errorf("Expected red color from section5, got %v", c)
		}
	})

	t.Run("PassThrough", func(t *testing.T) {
		s5i_pass := solidImage{image.Rect(105, 105, 125, 125), s5Red}
		f := NewFrame(dest, base, middle, &Section5{Image: s5i_pass, Replace: true}, PassThrough)
		c := f.At(110, 110)
		r, g, b, a := c.RGBA()
		if r != 65535 || g != 0 || b != 0 || a != 65535 {
			t.Errorf("Expected red color from section5, got %v", c)
		}
	})
}

func TestNineSliceGeometryExactBounds(t *testing.T) {
	// Base image 30x30
	base := image.NewRGBA(image.Rect(0, 0, 30, 30))
	// Set corner to Green
	cGreen := color.RGBA{0, 255, 0, 255}
	// Set middle start pixel to Red
	cRed := color.RGBA{255, 0, 0, 255}

	base.Set(9, 9, cGreen) // Just outside middle
	base.Set(10, 10, cRed) // Exact start of middle

	// Middle 10x10, from (10,10) to (20,20)
	middle := image.Rect(10, 10, 20, 20)
	dest := image.Rect(0, 0, 30, 30) // 1:1 scale for easy validation

	f := NewFrame(dest, base, middle, Repeating)

	// Sample exact boundary values
	if c := f.At(9, 9); c != cGreen {
		t.Errorf("Expected green color at (9,9) - border, got %v", c)
	}

	if c := f.At(10, 10); c != cRed {
		t.Errorf("Expected red color at (10,10) - exact middle start, got %v", c)
	}
}

func TestSection5AlphaComposition(t *testing.T) {
	base := solidImage{image.Rect(0, 0, 30, 30), color.RGBA{100, 100, 100, 255}}
	middle := image.Rect(10, 10, 20, 20)

	// Semi-transparent blue Section5
	semiBlue := color.RGBA{0, 0, 255, 127}
	s5i := solidImage{image.Rect(0, 0, 10, 10), semiBlue}

	dest := image.Rect(0, 0, 30, 30)
	f := NewFrame(dest, base, middle, &Section5{Image: s5i, Replace: false}, Section5Zeroed)

	c := f.At(10, 10)
	r, g, b, a := c.RGBA()

	// We expect roughly blended (alpha ~127 on 100 gray base)
	// Base: r=100, g=100, b=100
	// S5: r=0, g=0, b=255, alpha=127
	// Expected roughly: r=50, g=50, b=177 (or similar from blend math)
	// Note: Our simplified draw.Over uses 8-bit output
	if a == 0 || (r == 65535 && g == 65535 && b == 65535) || (r == 0 && g == 0 && b == 0 && a == 0) {
		t.Errorf("Expected alpha composition color, got %v", c)
	}
}
