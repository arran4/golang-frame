package main

import (
	"bytes"
	_ "embed"
	"github.com/arran4/golang-frame"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var (
	//go:embed "frame.png"
	baseImageData []byte
	//go:embed "window.png"
	titleImageData []byte
)

type SubImagable interface {
	draw.Image
	SubImage(r image.Rectangle) image.Image
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	base, err := png.Decode(bytes.NewReader(baseImageData))
	if err != nil {
		log.Panicf("Error with loading base file: %s", err)
	}
	ti, err := png.Decode(bytes.NewReader(titleImageData))
	if err != nil {
		log.Panicf("Error with loading base file: %s", err)
	}
	s5i := image.NewRGBA(image.Rect(0, 0, 100, 100))
	const sz = 50
	draw.Draw(s5i, s5i.Rect, SrcImageFunc(func(x, y int) color.Color {
		if (x/2)/sz%2 == 0 && (x/2)/sz == y/sz {
			return color.RGBA{0, 0, 127, 127}
		}
		return color.RGBA{}
	}), image.Pt(0, 0), draw.Src)
	c := image.NewRGBA(image.Rect(0, 0, 120*3, 120*3))
	Draw(c, ti, base, s5i)
	SaveFile(c)
}

type SrcImageFunc func(x int, y int) color.Color

func (s SrcImageFunc) Convert(c color.Color) color.Color {
	return c
}

func (s SrcImageFunc) ColorModel() color.Model {
	return s
}

func (s SrcImageFunc) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{-1e9, -1e9}, image.Point{1e9, 1e9}}
}

func (s SrcImageFunc) At(x, y int) color.Color {
	return s(x, y)
}

func Draw(i SubImagable, tibase, fibase, s5i image.Image) {
	titletext := "Alpha test"
	gr, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Panicf("font load error: %s", err)
	}
	grf := truetype.NewFace(gr, &truetype.Options{
		Size: 8,
		DPI:  180,
	})
	xp := 10
	yp := 10
	tidst := i.SubImage(image.Rect(xp, yp, xp+300, yp+40)).(SubImagable)
	frdst := i.SubImage(image.Rect(xp, yp+40, xp+300, yp+260+40)).(SubImagable)
	s1 := tibase.Bounds()
	s1.Max.Y = 42
	ti := frame.NewFrame(tidst.Bounds(), tibase.(SubImagable).SubImage(s1), image.Rect(16, 16, 47, 42))

	fr := frame.NewFrame(frdst.Bounds(), fibase.(SubImagable), image.Rect(11, 11, 111, 97), &frame.Section5{Image: s5i, Replace: false}, frame.Section5Zeroed, frame.Stretched)

	draw.Draw(tidst, tidst.Bounds(), ti, tidst.Bounds().Min, draw.Src)

	grfd := &font.Drawer{
		Dst:  tidst.SubImage(ti.MiddleRect()).(draw.Image),
		Src:  image.NewUniform(colornames.Blue),
		Face: grf,
	}
	ttb, _ := grfd.BoundString(titletext)
	grfd.Dot = grfd.Dot.Sub(ttb.Min).Add(fixed.P(ti.MiddleRect().Min.X, ti.MiddleRect().Min.Y))
	grfd.DrawString(titletext)

	draw.Draw(frdst, frdst.Bounds(), fr, frdst.Bounds().Min, draw.Src)
}

func SaveFile(i *image.RGBA) {
	_ = os.MkdirAll("images", 0755)
	fi, err := os.Create("images/sample5.png")
	if err != nil {
		log.Panicf("File create error: %v", err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Panicf("File close error: %s", err)
		}
	}()
	if err := png.Encode(fi, i); err != nil {
		log.Panicf("PNG encoding error: %s", err)
	}
}
