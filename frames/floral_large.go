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
	Middle: image.Rect(32, 32, 96, 96),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FloralLargePng))
	if err != nil { panic(err) }
	FloralLarge.Image = img
}
