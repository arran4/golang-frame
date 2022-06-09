package frame

import (
	"image"
	"image/color"
	"math"
)

type Frame struct {
	// Dest ination image size
	Dest image.Rectangle
	// Base image to use for the sections
	Base image.Image
	// The middle section of the Base image
	Middle image.Rectangle
	// The Border mode to use
	BorderMode BorderMode
	// Section5Override an image to replace section 5 with
	Section5Override *Section5
	// Section5Pos is how to handle Section5Override positioning
	Section5Pos Section5Positioning
}

// ColorModel Pass through to base image.
func (f *Frame) ColorModel() color.Model {
	return f.Base.ColorModel()
}

// Bounds end bounds of image
func (f *Frame) Bounds() image.Rectangle {
	return f.Dest
}

// At overrides the at functionality with our own multiplexer version
func (f *Frame) At(x, y int) color.Color {
	xp := x - f.Dest.Min.X
	yp := y - f.Dest.Min.Y
	xDistanceFromEnd := f.Dest.Dx() - xp
	midStartX := f.Middle.Min.X - f.Base.Bounds().Min.X
	midEndX := f.Dest.Dx() - (f.Base.Bounds().Max.X - f.Middle.Max.X)
	s5 := 0
	if xDistanceFromEnd <= f.Base.Bounds().Max.Sub(f.Middle.Max).X {
		xp = f.Base.Bounds().Dx() - xDistanceFromEnd
	} else if xp > midStartX {
		switch f.BorderMode {
		case Stretched:
			xp = midStartX + int(float64(xp-midStartX)/float64(midEndX-midStartX)*float64(f.Middle.Dx()))
		default:
			xp = midStartX + (xp-midStartX)%(f.Middle.Dx())
		}
		s5++
	}
	yDistanceFromEnd := f.Dest.Dy() - yp
	midStartY := f.Middle.Min.Y - f.Base.Bounds().Min.Y
	midEndY := f.Dest.Dy() - (f.Base.Bounds().Max.Y - f.Middle.Max.Y)
	if yDistanceFromEnd <= f.Base.Bounds().Max.Sub(f.Middle.Max).Y {
		yp = f.Base.Bounds().Dy() - yDistanceFromEnd
	} else if yp > midStartY {
		switch f.BorderMode {
		case Stretched:
			yp = midStartY + int(float64(yp-midStartY)/float64(midEndY-midStartY)*float64(f.Middle.Dy()))
		default:
			yp = midStartY + (yp-midStartY)%(f.Middle.Dy())
		}
		s5++
	}
	if f.Section5Override != nil && s5 == 2 {
		s5b := f.Section5Override.Bounds()
		xp = x - f.Dest.Min.X
		yp = y - f.Dest.Min.Y
		var c color.Color
		switch f.Section5Pos {
		case PassThrough:
			c = f.Section5Override.At(x, y)
		case Zerod:
			c = f.Section5Override.At(xp, yp)
		default:
			c = f.Section5Override.At(xp-midStartX-s5b.Min.X, yp-midStartY-s5b.Min.Y)
		}
		r1, g1, b1, a1 := c.RGBA()
		if a1 == math.MaxUint16 || f.Section5Override.Replace {
			return c
		} else if a1 != 0 {
			destc := f.Base.At(xp+f.Base.Bounds().Min.X, yp+f.Base.Bounds().Min.Y)
			destr, destg, destb, desta := destc.RGBA()
			a := (math.MaxUint16 - a1) * 0x101
			return color.RGBA{
				uint8((uint32(destr)*a/math.MaxUint16 + r1) >> 8),
				uint8((uint32(destg)*a/math.MaxUint16 + g1) >> 8),
				uint8((uint32(destb)*a/math.MaxUint16 + b1) >> 8),
				uint8((uint32(desta)*a/math.MaxUint16 + a1) >> 8),
			}
		}
	}
	return f.Base.At(xp+f.Base.Bounds().Min.X, yp+f.Base.Bounds().Min.Y)
}

// MiddleRect calculates the total space in the resulting image that section 5 consumes
func (f *Frame) MiddleRect() image.Rectangle {
	return image.Rectangle{
		Min: f.Dest.Min.Add(f.Middle.Min.Sub(f.Base.Bounds().Min)),
		Max: f.Dest.Max.Sub(f.Base.Bounds().Max.Sub(f.Middle.Max)),
	}
}

// Options is the interface for options on NewFrame
type Options interface {
	Option(f *Frame)
}

// Section5 is an optional image to replace section 5 with
type Section5 struct {
	image.Image
	// Replace If the new section 5 image should fully replace the base section 5 image
	Replace bool
}

// Option enables the use as a config option
func (s5 *Section5) Option(f *Frame) {
	f.Section5Override = s5
}

// BorderMode refers to the algorithm to use for filling the gaps produced by the variable size of section 5
type BorderMode int

// Option enables the use as a config option
func (bm BorderMode) Option(f *Frame) {
	f.BorderMode = bm
}

// Section5Positioning refers to the algorithm to use for filling section 5 with it's replacement image
type Section5Positioning int

// Option enables the use as a config option
func (bm Section5Positioning) Option(f *Frame) {
	f.Section5Pos = bm
}

const (
	// Section5Zeroed tells the draw algorithm to 0 position the section 5 image
	Section5Zeroed Section5Positioning = iota
	// Zerod uses the internal 0ed position
	Zerod
	// PassThrough tells the draw algorithm to pass through the position to section 5 without modification
	PassThrough
	// Repeating tells the draw algorithm to repeat sections 2,4,6, and 8 as required
	Repeating BorderMode = iota
	// Stretched tells the draw algorithm to stretch sections 2,4,6, and 8 proportionally
	Stretched
)

// NewFrame creates a new frame.
// `dest` is the destination size. Base in the image to scale.
// `base` is the base image to use for the sections
// `middle` is the middle section, and defines the section borders on the image
// `ops` options to apply (namely BorderMode)
func NewFrame(dest image.Rectangle, base image.Image, middle image.Rectangle, ops ...Options) *Frame {
	f := &Frame{
		Dest:       dest,
		Base:       base,
		Middle:     middle,
		BorderMode: Repeating,
	}
	for _, o := range ops {
		o.Option(f)
	}
	return f
}
