package draw

import (
	"image"
	"image/color"
	"testing"
)

func TestN9_At(t *testing.T) {
	// Create dummy images of 2x2
	mkImg := func(val uint8) image.Image {
		img := image.NewRGBA(image.Rect(0, 0, 2, 2))
		for y := 0; y < 2; y++ {
			for x := 0; x < 2; x++ {
				img.Set(x, y, color.Gray{Y: val})
			}
		}
		return img
	}

	n9 := &N9{
		TopLeft: mkImg(1), Top: mkImg(2), TopRight: mkImg(3),
		Left: mkImg(4), Center: mkImg(5), Right: mkImg(6),
		BottomLeft: mkImg(7), Bottom: mkImg(8), BottomRight: mkImg(9),
	}

	// Bounds check
	// Each is 2x2. Total size should be 6x6.
	b := n9.Bounds()
	if b.Dx() != 6 || b.Dy() != 6 {
		t.Errorf("Expected bounds 6x6, got %dx%d", b.Dx(), b.Dy())
	}

	// Test mapping
	tests := []struct {
		x, y int
		val  uint8
	}{
		{0, 0, 1}, {1, 0, 1}, // TopLeft
		{2, 0, 2}, {3, 0, 2}, // Top
		{4, 0, 3}, {5, 0, 3}, // TopRight

		{0, 2, 4}, {1, 2, 4}, // Left
		{2, 2, 5}, {3, 2, 5}, // Center
		{4, 2, 6}, {5, 2, 6}, // Right

		{0, 4, 7}, {1, 4, 7}, // BottomLeft
		{2, 4, 8}, {3, 4, 8}, // Bottom
		{4, 4, 9}, {5, 4, 9}, // BottomRight
	}

	for _, tc := range tests {
		c := n9.At(tc.x, tc.y)
		g, ok := c.(color.Gray)
		if !ok {
			// Convert if necessary
			r, g2, b, _ := c.RGBA()
			if r>>8 != uint32(tc.val) || g2>>8 != uint32(tc.val) || b>>8 != uint32(tc.val) {
				t.Errorf("At(%d, %d) = %v, want %d", tc.x, tc.y, c, tc.val)
			}
		} else {
			if g.Y != tc.val {
				t.Errorf("At(%d, %d) = %d, want %d", tc.x, tc.y, g.Y, tc.val)
			}
		}
	}
}
