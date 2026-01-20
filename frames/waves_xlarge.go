package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed waves_xlarge.png
var WavesXlargePng []byte

var WavesXlarge = &Definition{
	Name: "waves_xlarge",
	Middle: image.Rect(24, 24, 168, 168),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WavesXlargePng))
	if err != nil { panic(err) }
	WavesXlarge.Image = img
}
