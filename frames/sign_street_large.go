package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_street_large.png
var SignStreetLargePng []byte

var SignStreetLarge = &Definition{
	Name: "sign_street_large",
	Middle: image.Rect(16, 16, 112, 112),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignStreetLargePng))
	if err != nil { panic(err) }
	SignStreetLarge.Image = img
}
