package draw

import (
	"image/color"
	"math"
)

// Over is the alpha mode applied by draw.Draw() Probably inaccurate.
func Over(baseC color.Color, a1 uint32, r1 uint32, g1 uint32, b1 uint32) color.Color {
	baser, baseg, baseb, basea := baseC.RGBA()
	invA1 := math.MaxUint16 - a1
	c := color.RGBA{
		R: uint8((uint32(baser)*invA1/math.MaxUint16 + r1*a1) >> 8),
		G: uint8((uint32(baseg)*invA1/math.MaxUint16 + g1*a1) >> 8),
		B: uint8((uint32(baseb)*invA1/math.MaxUint16 + b1*a1) >> 8),
		A: uint8((uint32(basea)*invA1/math.MaxUint16 + a1) >> 8),
	}
	return &c
}
