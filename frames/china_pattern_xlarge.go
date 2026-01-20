package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed china_pattern_xlarge.png
var ChinaPatternXlargePng []byte

var ChinaPatternXlarge = &Definition{
	Name: "china_pattern_xlarge",
	Middle: image.Rect(144, 144, 240, 240),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ChinaPatternXlargePng))
	if err != nil { panic(err) }
	ChinaPatternXlarge.Image = img
}
