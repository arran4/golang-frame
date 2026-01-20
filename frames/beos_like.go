package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed beos_like.png
var BeosLikePng []byte

var BeosLike = &Definition{
	Name: "beos_like",
	Middle: image.Rect(4, 4, 28, 28),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(BeosLikePng))
	if err != nil { panic(err) }
	BeosLike.Image = img
}
