package frame

import (
	"image"
	"image/color"
	"testing"
)

// patternedImage is a helper to create an image with unique pixel colors based on coordinates.
// This allows verifying exact coordinate mapping.
type patternedImage struct {
	rect image.Rectangle
}

func (p patternedImage) ColorModel() color.Model { return color.RGBAModel }
func (p patternedImage) Bounds() image.Rectangle { return p.rect }
func (p patternedImage) At(x, y int) color.Color {
	if !image.Pt(x, y).In(p.rect) {
		return color.Transparent
	}
	return color.RGBA{uint8(x % 256), uint8(y % 256), 255, 255}
}

func TestNineSliceGeometryExactBounds(t *testing.T) {
	// 1. Boundary classification: Test with discriminating Section5 override, test X/Y independently
	base := patternedImage{image.Rect(10, 10, 50, 50)}
	middle := image.Rect(20, 20, 40, 40)
	dest := image.Rect(100, 100, 150, 150)

	// Create a Section5 replacement that has different colors from the base
	s5i := image.NewRGBA(image.Rect(0, 0, 30, 30))
	s5Red := color.RGBA{255, 0, 0, 255}
	for y := 0; y < 30; y++ {
		for x := 0; x < 30; x++ {
			s5i.Set(x, y, s5Red)
		}
	}

	// Test without fix: if we don't have the fix (i.e. we use original `frame.go`), the test will fail
	// because exact boundary will not be from Section 5.
	f := NewFrame(dest, base, middle, &Section5{Image: s5i, Replace: true}, Section5Zeroed)

	// Destination middle starts at (110, 110) because dest starts at 100 and left/top border is 10.

	// Helper to check color
	isRed := func(c color.Color) bool {
		r, g, b, a := c.RGBA()
		return r == 65535 && g == 0 && b == 0 && a == 65535
	}
	isPattern := func(c color.Color, x, y int) bool {
		r, g, b, a := c.RGBA()
		expectedR := uint32(uint8(x%256)) * 0x101
		expectedG := uint32(uint8(y%256)) * 0x101
		return r == expectedR && g == expectedG && b == 65535 && a == 65535
	}

	// Check exact boundaries (midStartX=110, midStartY=110)
	// X-only boundary
	c1 := f.At(110, 120)
	if !isRed(c1) {
		t.Errorf("At (110, 120) expected red Section5, got %v", c1)
	}

	c2 := f.At(109, 120)
	if !isPattern(c2, 19, 30) {
		t.Errorf("At (109, 120) expected pattern (border), got %v", c2)
	}

	// Y-only boundary
	c3 := f.At(120, 110)
	if !isRed(c3) {
		t.Errorf("At (120, 110) expected red Section5, got %v", c3)
	}

	c4 := f.At(120, 109)
	if !isPattern(c4, 30, 19) {
		t.Errorf("At (120, 109) expected pattern (border), got %v", c4)
	}

	// Intersection
	c5 := f.At(110, 110)
	if !isRed(c5) {
		t.Errorf("At (110, 110) expected red Section5, got %v", c5)
	}

	c6 := f.At(109, 109)
	if !isPattern(c6, 19, 19) {
		t.Errorf("At (109, 109) expected pattern (border), got %v", c6)
	}
}

func TestSection5Bounds(t *testing.T) {
	// 2. Origin translation independently of boundaries.
	// Sample strictly inside Section 5, use nonzero/unequal X/Y source origins and a coordinate-distinct pixel pattern
	base := patternedImage{image.Rect(10, 10, 50, 50)}
	middle := image.Rect(20, 20, 40, 40)
	dest := image.Rect(100, 100, 150, 150)

	// Patterned Section 5 with non-zero, unequal bounds
	s5i := patternedImage{image.Rect(5, 7, 35, 37)}

	// Helper to check s5 pattern
	isS5Pattern := func(c color.Color, x, y int) bool {
		r, g, b, a := c.RGBA()
		expectedR := uint32(uint8(x%256)) * 0x101
		expectedG := uint32(uint8(y%256)) * 0x101
		return r == expectedR && g == expectedG && b == 65535 && a == 65535
	}

	// Destination middle starts at (110, 110) and is 30x30 in size.
	// We sample at (120, 120), strictly inside Section 5.
	sampleX, sampleY := 120, 120

	t.Run("Section5Zeroed", func(t *testing.T) {
		f := NewFrame(dest, base, middle, &Section5{Image: s5i, Replace: true}, Section5Zeroed)
		c := f.At(sampleX, sampleY)
		// Destination point is (120, 120).
		// np is (20, 20).
		// midStartX, midStartY is (10, 10) relative to destination origin (100, 100).
		// Section5Zeroed uses np - midStart + s5b.Min = (20, 20) - (10, 10) + (5, 7) = (15, 17)
		if !isS5Pattern(c, 15, 17) {
			t.Errorf("Section5Zeroed: At (%d, %d) expected s5 pattern at (15, 17), got %v", sampleX, sampleY, c)
		}
	})

	t.Run("Zerod", func(t *testing.T) {
		f := NewFrame(dest, base, middle, &Section5{Image: s5i, Replace: true}, Zerod)
		c := f.At(sampleX, sampleY)
		// Dest point (120, 120) -> np = (20, 20).
		// Zerod uses np + s5b.Min = (20, 20) + (5, 7) = (25, 27)
		if !isS5Pattern(c, 25, 27) {
			t.Errorf("Zerod: At (%d, %d) expected s5 pattern at (25, 27), got %v", sampleX, sampleY, c)
		}
	})

	t.Run("PassThrough", func(t *testing.T) {
		s5iPassThrough := patternedImage{image.Rect(90, 90, 160, 160)}
		f := NewFrame(dest, base, middle, &Section5{Image: s5iPassThrough, Replace: true}, PassThrough)
		c := f.At(sampleX, sampleY)
		// PassThrough uses op (120, 120).
		if !isS5Pattern(c, 120, 120) {
			t.Errorf("PassThrough: At (%d, %d) expected s5 pattern at (120, 120), got %v", sampleX, sampleY, c)
		}
	})
}

func TestSection5AlphaComposition(t *testing.T) {
	// 3. Actual alpha regression. Use color.NRGBA{0,0,255,127} over opaque RGBA{100,100,100,255}
	// Note: NRGBA is non-premultiplied. The `Draw.Over` expects premultiplied alpha (as returned by `RGBA()`).
	// `color.NRGBA{0, 0, 255, 127}.RGBA()` will return correctly premultiplied values.

	baseImg := image.NewRGBA(image.Rect(10, 10, 50, 50))
	grayColor := color.RGBA{100, 100, 100, 255}
	for y := 10; y < 50; y++ {
		for x := 10; x < 50; x++ {
			baseImg.Set(x, y, grayColor)
		}
	}
	middle := image.Rect(20, 20, 40, 40)
	dest := image.Rect(100, 100, 150, 150)

	s5i := image.NewRGBA(image.Rect(0, 0, 30, 30))
	semiBlue := color.NRGBA{0, 0, 255, 127}
	for y := 0; y < 30; y++ {
		for x := 0; x < 30; x++ {
			s5i.Set(x, y, semiBlue)
		}
	}

	f := NewFrame(dest, baseImg, middle, &Section5{Image: s5i, Replace: false}, Section5Zeroed)

	// Sample in middle
	c := f.At(120, 120)
	r, g, b, a := c.RGBA() // 16-bit values

	// Convert to 8-bit for comparison
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	a8 := uint8(a >> 8)

	// Expected blended color ~RGBA{50, 50, 177, 255} based on repository's `draw.Over` logic
	if r8 != 50 || g8 != 50 || b8 != 177 || a8 != 255 {
		t.Errorf("Expected alpha blended color {50, 50, 177, 255}, got {%d, %d, %d, %d}", r8, g8, b8, a8)
	}

	// Test Replace: true
	fReplace := NewFrame(dest, baseImg, middle, &Section5{Image: s5i, Replace: true}, Section5Zeroed)
	cRep := fReplace.At(120, 120)
	rRep, gRep, bRep, aRep := cRep.RGBA() // 16-bit
	// Expected 16-bit for NRGBA{0,0,255,127}: R=0, G=0, B=32639 (127 * 257), A=32639 (127 * 257)
	expectedB := uint32(127 * 257)
	expectedA := uint32(127 * 257)
	if rRep != 0 || gRep != 0 || bRep != expectedB || aRep != expectedA {
		t.Errorf("Expected replace color {0, 0, %d, %d}, got {%d, %d, %d, %d}", expectedB, expectedA, rRep, gRep, bRep, aRep)
	}
}
