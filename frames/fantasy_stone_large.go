package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed fantasy_stone_large.png
var FantasyStoneLargePng []byte

var FantasyStoneLarge = &Definition{
	Name: "fantasy_stone_large",
	Middle: image.Rect(40, 40, 152, 152),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FantasyStoneLargePng))
	if err != nil { panic(err) }
	FantasyStoneLarge.Image = img
}
