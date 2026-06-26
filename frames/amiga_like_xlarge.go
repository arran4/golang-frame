package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed amiga_like_xlarge.png
var AmigaLikeXlargePng []byte

var AmigaLikeXlarge = &Definition{
	Name: "amiga_like_xlarge",
	Middle: image.Rect(18, 18, 78, 78),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(AmigaLikeXlargePng))
	if err != nil { panic(err) }
	AmigaLikeXlarge.Image = img
}
