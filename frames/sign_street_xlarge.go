package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_street_xlarge.png
var SignStreetXlargePng []byte

var SignStreetXlarge = &Definition{
	Name: "sign_street_xlarge",
	Middle: image.Rect(24, 24, 168, 168),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignStreetXlargePng))
	if err != nil { panic(err) }
	SignStreetXlarge.Image = img
}
