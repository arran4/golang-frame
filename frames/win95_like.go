package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed win95_like.png
var Win95LikePng []byte

var Win95Like = &Definition{
	Name: "win95_like",
	Middle: image.Rect(4, 4, 28, 28),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(Win95LikePng))
	if err != nil { panic(err) }
	Win95Like.Image = img
}
