package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_system_3_like_large.png
var MacSystem3LikeLargePng []byte

var MacSystem3LikeLarge = &Definition{
	Name: "mac_system_3_like_large",
	Middle: image.Rect(4, 44, 96, 96),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacSystem3LikeLargePng))
	if err != nil { panic(err) }
	MacSystem3LikeLarge.Image = img
}
