package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed stars_xlarge.png
var StarsXlargePng []byte

var StarsXlarge = &Definition{
	Name: "stars_xlarge",
	Middle: image.Rect(24, 24, 168, 168),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(StarsXlargePng))
	if err != nil { panic(err) }
	StarsXlarge.Image = img
}
