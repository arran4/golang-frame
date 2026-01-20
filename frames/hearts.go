package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed hearts.png
var HeartsPng []byte

var Hearts = &Definition{
	Name: "hearts",
	Middle: image.Rect(16, 16, 48, 48),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(HeartsPng))
	if err != nil { panic(err) }
	Hearts.Image = img
}
