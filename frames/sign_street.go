package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_street.png
var SignStreetPng []byte

var SignStreet = &Definition{
	Name: "sign_street",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignStreetPng))
	if err != nil { panic(err) }
	SignStreet.Image = img
}
