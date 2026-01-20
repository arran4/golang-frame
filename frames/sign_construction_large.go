package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_construction_large.png
var SignConstructionLargePng []byte

var SignConstructionLarge = &Definition{
	Name: "sign_construction_large",
	Middle: image.Rect(32, 32, 160, 160),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignConstructionLargePng))
	if err != nil { panic(err) }
	SignConstructionLarge.Image = img
}
