package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed fantasy_stone.png
var FantasyStonePng []byte

var FantasyStone = &Definition{
	Name: "fantasy_stone",
	Middle: image.Rect(24, 24, 72, 72),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FantasyStonePng))
	if err != nil { panic(err) }
	FantasyStone.Image = img
}
