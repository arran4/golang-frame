package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed metal_large.png
var MetalLargePng []byte

var MetalLarge = &Definition{
	Name: "metal_large",
	Middle: image.Rect(16, 16, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MetalLargePng))
	if err != nil { panic(err) }
	MetalLarge.Image = img
}
