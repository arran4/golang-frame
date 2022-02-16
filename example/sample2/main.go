package main

import (
	"bytes"
	_ "embed"
	"github.com/arran4/frame"
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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	base, err := png.Decode(bytes.NewReader(baseImageData))
	if err != nil {
		log.Panicf("Error with loading base file: %s", err)
	}
	i := image.NewRGBA(image.Rect(0, 0, 600, 600))
	dst := i.SubImage(image.Rect(100, 100, 400, 400)).(draw.Image)
	fr := frame.NewFrame(dst.Bounds(), base, image.Rect(48, 48, 55, 66))
	draw.Draw(dst, dst.Bounds(), fr, dst.Bounds().Min, draw.Src)
	SaveFile(i)
}

func SaveFile(i *image.RGBA) {
	os.MkdirAll("images", 0755)
	fi, err := os.Create("images/sample2.png")
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
