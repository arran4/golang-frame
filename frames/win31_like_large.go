package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win31_like_large.png
var Win31LikeLargePng []byte

var Win31LikeLarge = &Definition{
	Name: "win31_like_large",
	Middle: image.Rect(50, 50, 104, 180),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win31LikeLargePng))
	if err != nil { panic(err) }
	Win31LikeLarge.Image = img
}
