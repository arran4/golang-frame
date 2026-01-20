package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed next_like.png
var NextLikePng []byte

var NextLike = &Definition{
	Name: "next_like",
	Middle: image.Rect(6, 6, 26, 26),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(NextLikePng))
	if err != nil { panic(err) }
	NextLike.Image = img
}
