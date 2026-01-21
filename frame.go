package frame

import (
	"fmt"

	"github.com/arran4/golang-frame/draw"
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
	geom        *frameGeometry
}

type frameGeometry struct {
	destDx, destDy                                     int
	baseBoundsMinX, baseBoundsMinY                     int
	baseBoundsMaxSubMiddleMaxX, baseBoundsMaxSubMiddleMaxY int
	middleDx, middleDy                                 int
	midStartX, midStartY                               int
	midEndX, midEndY                                   int
	baseBoundsDx, baseBoundsDy                         int
}

// ColorModel Pass through to base image.
func (f *Frame) ColorModel() color.Model {
	return f.Base.ColorModel()
}

// Bounds end bounds of image
func (f *Frame) Bounds() image.Rectangle {
	return f.Dest
}

func (f *Frame) ensureGeom() {
	if f.geom != nil {
		return
	}
	f.geom = &frameGeometry{
		destDx:                     f.Dest.Dx(),
		destDy:                     f.Dest.Dy(),
		baseBoundsMinX:             f.Base.Bounds().Min.X,
		baseBoundsMinY:             f.Base.Bounds().Min.Y,
		baseBoundsMaxSubMiddleMaxX: f.Base.Bounds().Max.Sub(f.Middle.Max).X,
		baseBoundsMaxSubMiddleMaxY: f.Base.Bounds().Max.Sub(f.Middle.Max).Y,
		middleDx:                   f.Middle.Dx(),
		middleDy:                   f.Middle.Dy(),
		midStartX:                  f.Middle.Min.X - f.Base.Bounds().Min.X,
		midStartY:                  f.Middle.Min.Y - f.Base.Bounds().Min.Y,
		midEndX:                    f.Dest.Dx() - (f.Base.Bounds().Max.X - f.Middle.Max.X),
		midEndY:                    f.Dest.Dy() - (f.Base.Bounds().Max.Y - f.Middle.Max.Y),
		baseBoundsDx:               f.Base.Bounds().Dx(),
		baseBoundsDy:               f.Base.Bounds().Dy(),
	}
}

// At overrides the at functionality with our own multiplexer version
func (f *Frame) At(x, y int) color.Color {
	f.ensureGeom()
	op := image.Pt(x, y)
	p := op.Sub(f.Dest.Min)
	xDistanceFromEnd := f.geom.destDx - p.X
	s5 := 0
	if xDistanceFromEnd <= f.geom.baseBoundsMaxSubMiddleMaxX {
		p.X = f.geom.baseBoundsDx - xDistanceFromEnd
	} else if p.X > f.geom.midStartX {
		switch f.BorderMode {
		case Stretched:
			p.X = f.geom.midStartX + int(float64(p.X-f.geom.midStartX)/float64(f.geom.midEndX-f.geom.midStartX)*float64(f.geom.middleDx))
		default:
			p.X = f.geom.midStartX + (p.X-f.geom.midStartX)%(f.geom.middleDx)
		}
		s5++
	}
	yDistanceFromEnd := f.geom.destDy - p.Y
	if yDistanceFromEnd <= f.geom.baseBoundsMaxSubMiddleMaxY {
		p.Y = f.geom.baseBoundsDy - yDistanceFromEnd
	} else if p.Y > f.geom.midStartY {
		switch f.BorderMode {
		case Stretched:
			p.Y = f.geom.midStartY + int(float64(p.Y-f.geom.midStartY)/float64(f.geom.midEndY-f.geom.midStartY)*float64(f.geom.middleDy))
		default:
			p.Y = f.geom.midStartY + (p.Y-f.geom.midStartY)%(f.geom.middleDy)
		}
		s5++
	}
	if f.Section5Override != nil && s5 == 2 {
		if c2 := f.Section5Override.Apply(p, op, f, f.geom.midStartX, f.geom.midStartY); c2 != nil {
			return c2
		}
	}
	return f.Base.At(p.X+f.geom.baseBoundsMinX, p.Y+f.geom.baseBoundsMinY)
}

// Apply applies the section 5 override specific code.
func (s5 *Section5) Apply(p image.Point, op image.Point, f *Frame, midStartX int, midStartY int) color.Color {
	s5b := s5.Bounds()
	np := op.Sub(f.Dest.Min)
	var c color.Color
	switch f.Section5Pos {
	case PassThrough:
		c = s5.At(op.X, op.Y)
	case Zerod:
		c = s5.At(np.X, np.Y)
	default:
		c = s5.At(np.X-midStartX-s5b.Min.X, np.Y-midStartY-s5b.Min.Y)
	}
	if s5.Replace {
		return c
	}
	r1, g1, b1, a1 := c.RGBA()
	if a1 == math.MaxUint16 {
		return c
	} else if a1 == 0 {
		return nil
	}
	baseC := f.Base.At(p.X+f.Base.Bounds().Min.X, p.Y+f.Base.Bounds().Min.Y)
	if s5.AlphaMode == nil {
		return draw.Over(baseC, a1, r1, g1, b1)
	} else {
		return s5.AlphaMode(baseC, a1, r1, g1, b1)
	}
}

// MiddleRect calculates the total space in the resulting image that section 5 consumes
func (f *Frame) MiddleRect() image.Rectangle {
	f.ensureGeom()
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
	// AlphaMode the mode to apply alpha with, defaults to "over"
	AlphaMode func(destC color.Color, a1 uint32, r1 uint32, g1 uint32, b1 uint32) color.Color
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

// ValidateSize checks if the frame size is valid for the current configuration
func (f *Frame) ValidateSize(width, height int) error {
	if f.BorderMode != Repeating {
		return nil
	}
	// Width calculation:
	baseDx := f.Base.Bounds().Dx()
	midDx := f.Middle.Dx()
	if width < baseDx {
		return fmt.Errorf("width %d is too small, minimum is %d", width, baseDx)
	}
	if (width-baseDx)%midDx != 0 {
		lower := width - ((width - baseDx) % midDx)
		higher := lower + midDx
		return fmt.Errorf("width %d is invalid; try %d or %d", width, lower, higher)
	}
	// Height calculation:
	baseDy := f.Base.Bounds().Dy()
	midDy := f.Middle.Dy()
	if height < baseDy {
		return fmt.Errorf("height %d is too small, minimum is %d", height, baseDy)
	}
	if (height-baseDy)%midDy != 0 {
		lower := height - ((height - baseDy) % midDy)
		higher := lower + midDy
		return fmt.Errorf("height %d is invalid; try %d or %d", height, lower, higher)
	}
	return nil
}

// MeasureFrame returns the valid dimensions for a frame given the base image, middle rectangle, and target dimensions.
// It assumes Repeating border mode.
func MeasureFrame(base image.Image, middle image.Rectangle, targetWidth, targetHeight int) (widthLow, widthHigh, heightLow, heightHigh int) {
	baseDx := base.Bounds().Dx()
	midDx := middle.Dx()
	if targetWidth < baseDx {
		widthLow = baseDx
		widthHigh = baseDx
	} else {
		widthLow = targetWidth - ((targetWidth - baseDx) % midDx)
		widthHigh = widthLow + midDx
	}

	baseDy := base.Bounds().Dy()
	midDy := middle.Dy()
	if targetHeight < baseDy {
		heightLow = baseDy
		heightHigh = baseDy
	} else {
		heightLow = targetHeight - ((targetHeight - baseDy) % midDy)
		heightHigh = heightLow + midDy
	}
	return widthLow, widthHigh, heightLow, heightHigh
}
