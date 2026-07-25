package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_street_rounded_xlarge.png
var SignStreetRoundedXlargePng []byte

var SignStreetRoundedXlarge = &Definition{
	Name: "sign_street_rounded_xlarge",
	Middle: image.Rect(36, 36, 156, 156),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignStreetRoundedXlargePng))
	if err != nil { panic(err) }
	SignStreetRoundedXlarge.Image = img
}
