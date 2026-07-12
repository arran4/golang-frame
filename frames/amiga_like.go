package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed amiga_like.png
var AmigaLikePng []byte

var AmigaLike = &Definition{
	Name: "amiga_like",
	Middle: image.Rect(20, 13, 40, 57),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(AmigaLikePng))
	if err != nil { panic(err) }
	AmigaLike.Image = img
}
