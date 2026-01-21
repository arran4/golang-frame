package draw

import (
	"image"
	"image/color"
)

type N9 struct {
	TopLeft, Top, TopRight          image.Image
	Left, Center, Right             image.Image
	BottomLeft, Bottom, BottomRight image.Image
}

func (n *N9) ColorModel() color.Model {
	return color.RGBAModel
}

func (n *N9) Bounds() image.Rectangle {
	dx := func(i image.Image) int {
		if i == nil {
			return 0
		}
		return i.Bounds().Dx()
	}
	dy := func(i image.Image) int {
		if i == nil {
			return 0
		}
		return i.Bounds().Dy()
	}

	c0 := 0
	if v := dx(n.TopLeft); v > c0 {
		c0 = v
	}
	if v := dx(n.Left); v > c0 {
		c0 = v
	}
	if v := dx(n.BottomLeft); v > c0 {
		c0 = v
	}
	c1 := 0
	if v := dx(n.Top); v > c1 {
		c1 = v
	}
	if v := dx(n.Center); v > c1 {
		c1 = v
	}
	if v := dx(n.Bottom); v > c1 {
		c1 = v
	}
	c2 := 0
	if v := dx(n.TopRight); v > c2 {
		c2 = v
	}
	if v := dx(n.Right); v > c2 {
		c2 = v
	}
	if v := dx(n.BottomRight); v > c2 {
		c2 = v
	}

	r0 := 0
	if v := dy(n.TopLeft); v > r0 {
		r0 = v
	}
	if v := dy(n.Top); v > r0 {
		r0 = v
	}
	if v := dy(n.TopRight); v > r0 {
		r0 = v
	}
	r1 := 0
	if v := dy(n.Left); v > r1 {
		r1 = v
	}
	if v := dy(n.Center); v > r1 {
		r1 = v
	}
	if v := dy(n.Right); v > r1 {
		r1 = v
	}
	r2 := 0
	if v := dy(n.BottomLeft); v > r2 {
		r2 = v
	}
	if v := dy(n.Bottom); v > r2 {
		r2 = v
	}
	if v := dy(n.BottomRight); v > r2 {
		r2 = v
	}

	return image.Rect(0, 0, c0+c1+c2, r0+r1+r2)
}

func (n *N9) At(x, y int) color.Color {
	dx := func(i image.Image) int {
		if i == nil {
			return 0
		}
		return i.Bounds().Dx()
	}
	dy := func(i image.Image) int {
		if i == nil {
			return 0
		}
		return i.Bounds().Dy()
	}

	c0 := 0
	if v := dx(n.TopLeft); v > c0 {
		c0 = v
	}
	if v := dx(n.Left); v > c0 {
		c0 = v
	}
	if v := dx(n.BottomLeft); v > c0 {
		c0 = v
	}
	c1 := 0
	if v := dx(n.Top); v > c1 {
		c1 = v
	}
	if v := dx(n.Center); v > c1 {
		c1 = v
	}
	if v := dx(n.Bottom); v > c1 {
		c1 = v
	}

	r0 := 0
	if v := dy(n.TopLeft); v > r0 {
		r0 = v
	}
	if v := dy(n.Top); v > r0 {
		r0 = v
	}
	if v := dy(n.TopRight); v > r0 {
		r0 = v
	}
	r1 := 0
	if v := dy(n.Left); v > r1 {
		r1 = v
	}
	if v := dy(n.Center); v > r1 {
		r1 = v
	}
	if v := dy(n.Right); v > r1 {
		r1 = v
	}

	var target image.Image
	localX, localY := x, y

	if y < r0 {
		// Top row
		if x < c0 {
			target = n.TopLeft
		} else if x < c0+c1 {
			target = n.Top
			localX -= c0
		} else {
			target = n.TopRight
			localX -= (c0 + c1)
		}
	} else if y < r0+r1 {
		// Middle row
		localY -= r0
		if x < c0 {
			target = n.Left
		} else if x < c0+c1 {
			target = n.Center
			localX -= c0
		} else {
			target = n.Right
			localX -= (c0 + c1)
		}
	} else {
		// Bottom row
		localY -= (r0 + r1)
		if x < c0 {
			target = n.BottomLeft
		} else if x < c0+c1 {
			target = n.Bottom
			localX -= c0
		} else {
			target = n.BottomRight
			localX -= (c0 + c1)
		}
	}

	if target == nil {
		return color.RGBA{}
	}

	// Adjust for image bounds Min (usually 0, but good to be safe)
	return target.At(target.Bounds().Min.X+localX, target.Bounds().Min.Y+localY)
}
