package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed scifi_tech.png
var ScifiTechPng []byte

var ScifiTech = &Definition{
	Name: "scifi_tech",
	Middle: image.Rect(24, 24, 72, 72),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ScifiTechPng))
	if err != nil { panic(err) }
	ScifiTech.Image = img
}
