package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed hearts_large.png
var HeartsLargePng []byte

var HeartsLarge = &Definition{
	Name: "hearts_large",
	Middle: image.Rect(32, 32, 96, 96),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(HeartsLargePng))
	if err != nil { panic(err) }
	HeartsLarge.Image = img
}
