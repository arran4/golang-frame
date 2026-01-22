package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed gold_large.png
var GoldLargePng []byte

var GoldLarge = &Definition{
	Name: "gold_large",
	Middle: image.Rect(32, 32, 160, 160),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(GoldLargePng))
	if err != nil { panic(err) }
	GoldLarge.Image = img
}
