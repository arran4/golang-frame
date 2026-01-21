package frame

import (
	"image"
	"testing"
)

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
