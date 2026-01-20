package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed wood_xlarge.png
var WoodXlargePng []byte

var WoodXlarge = &Definition{
	Name: "wood_xlarge",
	Middle: image.Rect(48, 48, 240, 240),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WoodXlargePng))
	if err != nil { panic(err) }
	WoodXlarge.Image = img
}
