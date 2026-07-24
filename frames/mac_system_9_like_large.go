package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_system_9_like_large.png
var MacSystem9LikeLargePng []byte

var MacSystem9LikeLarge = &Definition{
	Name: "mac_system_9_like_large",
	Middle: image.Rect(6, 46, 94, 94),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacSystem9LikeLargePng))
	if err != nil { panic(err) }
	MacSystem9LikeLarge.Image = img
}
