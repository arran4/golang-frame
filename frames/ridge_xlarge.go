package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed ridge_xlarge.png
var RidgeXlargePng []byte

var RidgeXlarge = &Definition{
	Name: "ridge_xlarge",
	Middle: image.Rect(24, 24, 120, 120),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(RidgeXlargePng))
	if err != nil { panic(err) }
	RidgeXlarge.Image = img
}
