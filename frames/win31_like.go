package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win31_like.png
var Win31LikePng []byte

var Win31Like = &Definition{
	Name: "win31_like",
	Middle: image.Rect(4, 4, 28, 28),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win31LikePng))
	if err != nil { panic(err) }
	Win31Like.Image = img
}
