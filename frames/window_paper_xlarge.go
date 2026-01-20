package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_paper_xlarge.png
var WindowPaperXlargePng []byte

var WindowPaperXlarge = &Definition{
	Name: "window_paper_xlarge",
	Middle: image.Rect(24, 48, 156, 156),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowPaperXlargePng))
	if err != nil { panic(err) }
	WindowPaperXlarge.Image = img
}
