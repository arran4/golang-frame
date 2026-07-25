package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_street_rounded.png
var SignStreetRoundedPng []byte

var SignStreetRounded = &Definition{
	Name: "sign_street_rounded",
	Middle: image.Rect(12, 12, 52, 52),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignStreetRoundedPng))
	if err != nil { panic(err) }
	SignStreetRounded.Image = img
}
