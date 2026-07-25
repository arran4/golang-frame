package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_street_rounded_large.png
var SignStreetRoundedLargePng []byte

var SignStreetRoundedLarge = &Definition{
	Name: "sign_street_rounded_large",
	Middle: image.Rect(24, 24, 104, 104),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignStreetRoundedLargePng))
	if err != nil { panic(err) }
	SignStreetRoundedLarge.Image = img
}
