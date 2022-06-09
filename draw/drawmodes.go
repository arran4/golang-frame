package draw

import (
	"image/color"
	"math"
)

// Over is the alpha mode applied by draw.Draw() Probably inaccurate.
func Over(destc color.Color, a1 uint32, r1 uint32, g1 uint32, b1 uint32) color.Color {
	destr, destg, destb, desta := destc.RGBA()
	reva := (math.MaxUint16 - a1) * 0x101
	c := color.RGBA{
		R: uint8((uint32(destr)*reva/math.MaxUint16 + r1*a1) >> 8),
		G: uint8((uint32(destg)*reva/math.MaxUint16 + g1*a1) >> 8),
		B: uint8((uint32(destb)*reva/math.MaxUint16 + b1*a1) >> 8),
		A: uint8((uint32(desta)*reva/math.MaxUint16 + a1) >> 8),
	}
	return &c
}
