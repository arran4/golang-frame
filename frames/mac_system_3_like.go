package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_system_3_like.png
var MacSystem3LikePng []byte

var MacSystem3Like = &Definition{
	Name: "mac_system_3_like",
	Middle: image.Rect(2, 22, 48, 48),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacSystem3LikePng))
	if err != nil { panic(err) }
	MacSystem3Like.Image = img
}
