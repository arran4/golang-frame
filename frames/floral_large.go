package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed floral_large.png
var FloralLargePng []byte

var FloralLarge = &Definition{
	Name: "floral_large",
	Middle: image.Rect(64, 64, 128, 128),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FloralLargePng))
	if err != nil { panic(err) }
	FloralLarge.Image = img
}
