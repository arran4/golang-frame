package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_paper.png
var WindowPaperPng []byte

var WindowPaper = &Definition{
	Name: "window_paper",
	Middle: image.Rect(8, 16, 52, 52),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowPaperPng))
	if err != nil { panic(err) }
	WindowPaper.Image = img
}
