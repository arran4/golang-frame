package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed china_pattern_large.png
var ChinaPatternLargePng []byte

var ChinaPatternLarge = &Definition{
	Name: "china_pattern_large",
	Middle: image.Rect(96, 96, 160, 160),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ChinaPatternLargePng))
	if err != nil { panic(err) }
	ChinaPatternLarge.Image = img
}
