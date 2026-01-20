package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win95_like_large.png
var Win95LikeLargePng []byte

var Win95LikeLarge = &Definition{
	Name: "win95_like_large",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win95LikeLargePng))
	if err != nil { panic(err) }
	Win95LikeLarge.Image = img
}
