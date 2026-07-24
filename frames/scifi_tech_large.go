package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed scifi_tech_large.png
var ScifiTechLargePng []byte

var ScifiTechLarge = &Definition{
	Name: "scifi_tech_large",
	Middle: image.Rect(48, 48, 144, 144),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ScifiTechLargePng))
	if err != nil { panic(err) }
	ScifiTechLarge.Image = img
}
