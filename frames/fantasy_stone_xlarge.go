package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed fantasy_stone_xlarge.png
var FantasyStoneXlargePng []byte

var FantasyStoneXlarge = &Definition{
	Name: "fantasy_stone_xlarge",
	Middle: image.Rect(60, 60, 228, 228),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(FantasyStoneXlargePng))
	if err != nil { panic(err) }
	FantasyStoneXlarge.Image = img
}
