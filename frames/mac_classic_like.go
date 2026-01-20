package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_classic_like.png
var MacClassicLikePng []byte

var MacClassicLike = &Definition{
	Name: "mac_classic_like",
	Middle: image.Rect(4, 4, 28, 28),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacClassicLikePng))
	if err != nil { panic(err) }
	MacClassicLike.Image = img
}
