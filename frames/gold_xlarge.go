package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed gold_xlarge.png
var GoldXlargePng []byte

var GoldXlarge = &Definition{
	Name: "gold_xlarge",
	Middle: image.Rect(48, 48, 240, 240),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(GoldXlargePng))
	if err != nil { panic(err) }
	GoldXlarge.Image = img
}
