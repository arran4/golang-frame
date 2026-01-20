package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed waves_large.png
var WavesLargePng []byte

var WavesLarge = &Definition{
	Name: "waves_large",
	Middle: image.Rect(16, 16, 112, 112),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WavesLargePng))
	if err != nil { panic(err) }
	WavesLarge.Image = img
}
