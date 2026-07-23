package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_system_9_like.png
var MacSystem9LikePng []byte

var MacSystem9Like = &Definition{
	Name: "mac_system_9_like",
	Middle: image.Rect(3, 22, 48, 48),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacSystem9LikePng))
	if err != nil { panic(err) }
	MacSystem9Like.Image = img
}
