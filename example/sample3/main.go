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
	"image/draw"
	"image/png"
	"log"
	"os"
)

var (
	//go:embed "window.png"
	baseImageData []byte
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
	i := image.NewRGBA(image.Rect(0, 0, 630, 320))
	WindowRepeating(i, base)
	WindowStretched(i, base)
	SaveFile(i)
}

func WindowRepeating(i SubImagable, base image.Image) {
	titletext := "Repeating (default)"
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
	s1 := base.Bounds()
	s1.Max.Y = 42
	s2 := base.Bounds().Add(image.Pt(0, 42))
	s2.Max.Y += -42
	ti := frame.NewFrame(tidst.Bounds(), base.(SubImagable).SubImage(s1), image.Rect(16, 16, 47, 42))
	fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66))

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

func WindowStretched(i SubImagable, base image.Image) {
	titletext := "Stretched (Optional)"
	gr, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Panicf("font load error: %s", err)
	}
	grf := truetype.NewFace(gr, &truetype.Options{
		Size: 8,
		DPI:  180,
	})
	xp := 320
	yp := 10
	tidst := i.SubImage(image.Rect(xp, yp, xp+300, yp+40)).(SubImagable)
	frdst := i.SubImage(image.Rect(xp, yp+40, xp+300, yp+260+40)).(SubImagable)
	s1 := base.Bounds()
	s1.Max.Y = 42
	s2 := base.Bounds().Add(image.Pt(0, 42))
	s2.Max.Y += -42
	ti := frame.NewFrame(tidst.Bounds(), base.(SubImagable).SubImage(s1), image.Rect(16, 16, 47, 42), frame.Stretched)
	fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66), frame.Stretched)

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
	fi, err := os.Create("images/sample3.png")
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
