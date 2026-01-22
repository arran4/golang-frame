package frame

import (
	"image"
	"image/color"
	"testing"
)

func BenchmarkFrameAt_Stretched(b *testing.B) {
	// Setup a frame with Stretched border mode
	base := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// Fill base with unique colors to verify mapping if needed,
	// but for benchmark we just want to execute the path.
	// Let's just use 0,0,0,0 for speed of setup, or simple pattern.

	middle := image.Rect(10, 10, 90, 90)
	dest := image.Rect(0, 0, 200, 200)

	f := NewFrame(dest, base, middle, Stretched)

	// We want to hit the optimized path.
	// The path is hit when p.X > midStartX (left border end)
	// AND p.X < destDx - right_border_width (implied by first if/else logic)
	// Similar for Y.

	// midStartX = 10 - 0 = 10.
	// right border width = 100 - 90 = 10.
	// destDx = 200.
	// right border start in dest = 200 - 10 = 190.
	// So X in [11, 189] hits the stretch logic.

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Sample a few points in the stretched region
		f.At(50, 50)
		f.At(100, 100)
		f.At(150, 150)
		// Sample mixed axes (stretch X, regular Y? No, At checks X and Y independently but sequentially)
		f.At(50, 5) // X stretched, Y top border
	}
}

func TestFrameAt_Stretched_Correctness(t *testing.T) {
	// We verify that the mapping produces expected coordinates (or consistent ones).
	// Since we are changing implementation, we might want to capture current behavior
	// or ensure logical correctness.

	base := image.NewRGBA(image.Rect(0, 0, 10, 10))
	// 3x3 middle: (1,1) to (4,4) -> width 3, height 3.
	// indices 1, 2, 3 are middle.
	middle := image.Rect(1, 1, 4, 4)

	// Dest: 20x20.
	dest := image.Rect(0, 0, 20, 20)

	f := NewFrame(dest, base, middle, Stretched)

	// midStartX = 1.
	// midEndX = 20 - (10 - 4) = 20 - 6 = 14.
	// Dest middle width = 14 - 1 = 13.
	// Source middle width = 3.

	// Point x=7 (approx middle of dest middle region).
	// Formula: start + (7 - 1) * 3 / 13 = 1 + 6 * 3 / 13 = 1 + 18/13 = 1 + 1 = 2.
	// Float: 1 + 6 * (3/13) = 1 + 6 * 0.2307 = 1 + 1.38 = 2.38 -> int(2.38) = 2.

	// Point x=13 (near end).
	// Formula: 1 + (13 - 1) * 3 / 13 = 1 + 12 * 3 / 13 = 1 + 36/13 = 1 + 2 = 3.
	// Float: 1 + 12 * 0.2307 = 1 + 2.76 = 3.76 -> int(3.76) = 3.

	// So for these inputs, results should be stable.

	// We can check specific pixel coordinates mapping.
	// However, At returns Color. We need to check what pixel from Base is picked.
	// Let's set distinct colors in Base.

	// (1,1) -> Red
	// (2,1) -> Green
	// (3,1) -> Blue
	cRed := color.RGBA{255, 0, 0, 255}
	cGreen := color.RGBA{0, 255, 0, 255}
	cBlue := color.RGBA{0, 0, 255, 255}

	base.Set(1, 1, cRed)
	base.Set(2, 1, cGreen)
	base.Set(3, 1, cBlue)

	// We'll check X mapping at Y=1 (top of middle, no vertical stretch/offset if Y is handled correctly)
	// Y=1 in dest should map to Y=1 in base (since top border is 1px, same).

	// Dest X=7 -> expects Base X=2 (Green)
	if c := f.At(7, 1); c != cGreen {
		t.Errorf("At(7,1) = %v, want Green", c)
	}

	// Dest X=13 -> expects Base X=3 (Blue)
	if c := f.At(13, 1); c != cBlue {
		t.Errorf("At(13,1) = %v, want Blue", c)
	}

	// Dest X=1 -> Start of middle.
	// 1 + 0 = 1. Base X=1 (Red)
	if c := f.At(1, 1); c != cRed {
		t.Errorf("At(1,1) = %v, want Red", c)
	}
}

func TestValidateSize(t *testing.T) {
	// Create a dummy base image and middle rect
	base := image.NewRGBA(image.Rect(0, 0, 50, 50))
	middle := image.Rect(10, 10, 40, 40) // Middle is 30x30.
	// Left border = 10, Right border = 10.
	// Base width = 50.
	// Middle width = 30.

	f := &Frame{
		Base:       base,
		Middle:     middle,
		BorderMode: Repeating,
	}

	// Case 1: Base width (50) -> Valid
	if err := f.ValidateSize(50, 50); err != nil {
		t.Errorf("Expected size 50x50 to be valid, got: %v", err)
	}

	// Case 2: Base width + middle width (80) -> Valid
	if err := f.ValidateSize(80, 80); err != nil {
		t.Errorf("Expected size 80x80 to be valid, got: %v", err)
	}

	// Case 3: 60 (Base + 10) -> Invalid. Should suggest 50 and 80.
	err := f.ValidateSize(60, 60)
	if err == nil {
		t.Errorf("Expected size 60x60 to be invalid")
	} else {
		msg := err.Error()
		if msg != "width 60 is invalid; try 50 or 80" {
			t.Errorf("Unexpected error message: %s", msg)
		}
	}

	// Case 4: Too small
	err = f.ValidateSize(40, 40)
	if err == nil {
		t.Errorf("Expected size 40x40 to be invalid (too small)")
	}
}

func TestMeasureFrame(t *testing.T) {
	base := image.NewRGBA(image.Rect(0, 0, 50, 50))
	middle := image.Rect(10, 10, 40, 40) // Middle 30x30

	// Target 60. Base 50. Mid 30.
	// 60 - 50 = 10. 10 % 30 != 0.
	// Lower = 60 - 10 = 50.
	// Higher = 50 + 30 = 80.
	wl, wh, hl, hh := MeasureFrame(base, middle, 60, 60)
	if wl != 50 || wh != 80 {
		t.Errorf("Expected width 50, 80. Got %d, %d", wl, wh)
	}
	if hl != 50 || hh != 80 {
		t.Errorf("Expected height 50, 80. Got %d, %d", hl, hh)
	}

	// Target 40 (Too small). Base 50.
	wl, wh, hl, hh = MeasureFrame(base, middle, 40, 40)
	if wl != 50 || wh != 50 {
		t.Errorf("Expected width 50, 50. Got %d, %d", wl, wh)
	}
}
