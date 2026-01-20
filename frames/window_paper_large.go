package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_paper_large.png
var WindowPaperLargePng []byte

var WindowPaperLarge = &Definition{
	Name: "window_paper_large",
	Middle: image.Rect(16, 32, 104, 104),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowPaperLargePng))
	if err != nil { panic(err) }
	WindowPaperLarge.Image = img
}
