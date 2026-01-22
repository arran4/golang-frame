package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed next_like_large.png
var NextLikeLargePng []byte

var NextLikeLarge = &Definition{
	Name: "next_like_large",
	Middle: image.Rect(10, 36, 86, 86),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(NextLikeLargePng))
	if err != nil { panic(err) }
	NextLikeLarge.Image = img
}
