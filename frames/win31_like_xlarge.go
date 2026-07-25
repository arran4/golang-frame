package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win31_like_xlarge.png
var Win31LikeXlargePng []byte

var Win31LikeXlarge = &Definition{
	Name: "win31_like_xlarge",
	Middle: image.Rect(75, 75, 156, 270),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win31LikeXlargePng))
	if err != nil { panic(err) }
	Win31LikeXlarge.Image = img
}
