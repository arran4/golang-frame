package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed floral_xlarge.png
var FloralXlargePng []byte

var FloralXlarge = &Definition{
	Name: "floral_xlarge",
	Middle: image.Rect(96, 96, 192, 192),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FloralXlargePng))
	if err != nil { panic(err) }
	FloralXlarge.Image = img
}
