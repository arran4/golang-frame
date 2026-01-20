package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_construction_xlarge.png
var SignConstructionXlargePng []byte

var SignConstructionXlarge = &Definition{
	Name: "sign_construction_xlarge",
	Middle: image.Rect(48, 48, 240, 240),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignConstructionXlargePng))
	if err != nil { panic(err) }
	SignConstructionXlarge.Image = img
}
