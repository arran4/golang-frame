package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed mac_system_9_like_xlarge.png
var MacSystem9LikeXlargePng []byte

var MacSystem9LikeXlarge = &Definition{
	Name: "mac_system_9_like_xlarge",
	Middle: image.Rect(9, 69, 141, 141),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(MacSystem9LikeXlargePng))
	if err != nil { panic(err) }
	MacSystem9LikeXlarge.Image = img
}
