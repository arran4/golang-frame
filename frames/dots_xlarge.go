package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed dots_xlarge.png
var DotsXlargePng []byte

var DotsXlarge = &Definition{
	Name: "dots_xlarge",
	Middle: image.Rect(48, 48, 96, 96),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(DotsXlargePng))
	if err != nil { panic(err) }
	DotsXlarge.Image = img
}
