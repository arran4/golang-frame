package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed waves.png
var WavesPng []byte

var Waves = &Definition{
	Name: "waves",
	Middle: image.Rect(8, 8, 56, 56),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WavesPng))
	if err != nil { panic(err) }
	Waves.Image = img
}
