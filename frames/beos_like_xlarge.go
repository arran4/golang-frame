package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed beos_like_xlarge.png
var BeosLikeXlargePng []byte

var BeosLikeXlarge = &Definition{
	Name: "beos_like_xlarge",
	Middle: image.Rect(12, 12, 84, 84),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(BeosLikeXlargePng))
	if err != nil { panic(err) }
	BeosLikeXlarge.Image = img
}
