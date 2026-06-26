package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed chains_large.png
var ChainsLargePng []byte

var ChainsLarge = &Definition{
	Name: "chains_large",
	Middle: image.Rect(16, 16, 112, 112),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ChainsLargePng))
	if err != nil { panic(err) }
	ChainsLarge.Image = img
}
