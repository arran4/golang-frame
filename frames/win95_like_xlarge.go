package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win95_like_xlarge.png
var Win95LikeXlargePng []byte

var Win95LikeXlarge = &Definition{
	Name: "win95_like_xlarge",
	Middle: image.Rect(12, 12, 84, 84),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win95LikeXlargePng))
	if err != nil { panic(err) }
	Win95LikeXlarge.Image = img
}
