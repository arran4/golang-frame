package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed hearts_xlarge.png
var HeartsXlargePng []byte

var HeartsXlarge = &Definition{
	Name: "hearts_xlarge",
	Middle: image.Rect(48, 48, 144, 144),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(HeartsXlargePng))
	if err != nil { panic(err) }
	HeartsXlarge.Image = img
}
