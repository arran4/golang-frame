package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed next_like_xlarge.png
var NextLikeXlargePng []byte

var NextLikeXlarge = &Definition{
	Name: "next_like_xlarge",
	Middle: image.Rect(15, 54, 129, 129),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(NextLikeXlargePng))
	if err != nil { panic(err) }
	NextLikeXlarge.Image = img
}
