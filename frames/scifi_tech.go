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
	Middle: image.Rect(16, 16, 80, 80),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ScifiTechPng))
	if err != nil { panic(err) }
	ScifiTech.Image = img
}
