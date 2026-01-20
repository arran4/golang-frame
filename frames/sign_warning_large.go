package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_warning_large.png
var SignWarningLargePng []byte

var SignWarningLarge = &Definition{
	Name: "sign_warning_large",
	Middle: image.Rect(24, 24, 104, 104),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignWarningLargePng))
	if err != nil { panic(err) }
	SignWarningLarge.Image = img
}
