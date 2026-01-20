package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed scifi_tech_xlarge.png
var ScifiTechXlargePng []byte

var ScifiTechXlarge = &Definition{
	Name: "scifi_tech_xlarge",
	Middle: image.Rect(48, 48, 240, 240),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(ScifiTechXlargePng))
	if err != nil { panic(err) }
	ScifiTechXlarge.Image = img
}
