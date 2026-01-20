package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed stars.png
var StarsPng []byte

var Stars = &Definition{
	Name: "stars",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(StarsPng))
	if err != nil { panic(err) }
	Stars.Image = img
}
