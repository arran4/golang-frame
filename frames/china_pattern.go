package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed china_pattern.png
var ChinaPatternPng []byte

var ChinaPattern = &Definition{
	Name: "china_pattern",
	Middle: image.Rect(48, 48, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ChinaPatternPng))
	if err != nil { panic(err) }
	ChinaPattern.Image = img
}
