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
	Middle: image.Rect(18, 18, 78, 78),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(NextLikeXlargePng))
	if err != nil { panic(err) }
	NextLikeXlarge.Image = img
}
