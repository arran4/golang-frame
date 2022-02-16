package main

import (
	"github.com/arran4/frame"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func NewBasicFrame(r image.Rectangle) *frame.Frame {
	middle := image.Rect(2, 2, 3, 3)
	base := image.NewRGBA(image.Rect(0, 0, 5, 5))
	b := base.Bounds()
	for y, r := range [][]color.RGBA{
		{colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray},
		{colornames.Lightgray, colornames.Darkgrey, colornames.Darkgrey, colornames.Darkgrey, colornames.Lightgray},
		{colornames.Lightgray, colornames.Darkgrey, colornames.White, colornames.Darkgrey, colornames.Lightgray},
		{colornames.Lightgray, colornames.Darkgrey, colornames.Darkgrey, colornames.Darkgrey, colornames.Lightgray},
		{colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray},
	} {
		for x, c := range r {
			base.Set(b.Min.X+x, b.Min.Y+y, c)
		}
	}
	return frame.NewFrame(r, base, middle)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	i := image.NewRGBA(image.Rect(0, 0, 150, 100))
	targetArea := image.Rect(10, 10, 100, 30)
	fr := NewBasicFrame(targetArea)
	dst := i.SubImage(targetArea).(draw.Image)
	draw.Draw(dst, dst.Bounds(), fr, dst.Bounds().Min, draw.Src)
	SaveFile(i)
}

func SaveFile(i *image.RGBA) {
	os.MkdirAll("images", 0755)
	fi, err := os.Create("images/sample1.png")
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
