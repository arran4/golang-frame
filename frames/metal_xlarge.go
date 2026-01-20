package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed metal_xlarge.png
var MetalXlargePng []byte

var MetalXlarge = &Definition{
	Name: "metal_xlarge",
	Middle: image.Rect(24, 24, 120, 120),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MetalXlargePng))
	if err != nil { panic(err) }
	MetalXlarge.Image = img
}
