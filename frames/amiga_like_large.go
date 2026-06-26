package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed amiga_like_large.png
var AmigaLikeLargePng []byte

var AmigaLikeLarge = &Definition{
	Name: "amiga_like_large",
	Middle: image.Rect(12, 12, 52, 52),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(AmigaLikeLargePng))
	if err != nil { panic(err) }
	AmigaLikeLarge.Image = img
}
