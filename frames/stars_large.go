package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed stars_large.png
var StarsLargePng []byte

var StarsLarge = &Definition{
	Name: "stars_large",
	Middle: image.Rect(16, 16, 112, 112),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(StarsLargePng))
	if err != nil { panic(err) }
	StarsLarge.Image = img
}
