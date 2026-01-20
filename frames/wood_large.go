package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed wood_large.png
var WoodLargePng []byte

var WoodLarge = &Definition{
	Name: "wood_large",
	Middle: image.Rect(32, 32, 160, 160),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WoodLargePng))
	if err != nil { panic(err) }
	WoodLarge.Image = img
}
