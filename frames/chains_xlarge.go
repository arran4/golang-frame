package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed chains_xlarge.png
var ChainsXlargePng []byte

var ChainsXlarge = &Definition{
	Name: "chains_xlarge",
	Middle: image.Rect(52, 52, 236, 236),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ChainsXlargePng))
	if err != nil { panic(err) }
	ChainsXlarge.Image = img
}
