package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed chains.png
var ChainsPng []byte

var Chains = &Definition{
	Name: "chains",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ChainsPng))
	if err != nil { panic(err) }
	Chains.Image = img
}
