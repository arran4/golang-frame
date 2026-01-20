package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed ridge.png
var RidgePng []byte

var Ridge = &Definition{
	Name: "ridge",
	Middle: image.Rect(8, 8, 40, 40),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(RidgePng))
	if err != nil { panic(err) }
	Ridge.Image = img
}
