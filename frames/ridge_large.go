package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed ridge_large.png
var RidgeLargePng []byte

var RidgeLarge = &Definition{
	Name: "ridge_large",
	Middle: image.Rect(16, 16, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(RidgeLargePng))
	if err != nil { panic(err) }
	RidgeLarge.Image = img
}
