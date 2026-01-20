package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_warning.png
var SignWarningPng []byte

var SignWarning = &Definition{
	Name: "sign_warning",
	Middle: image.Rect(12, 12, 52, 52),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignWarningPng))
	if err != nil { panic(err) }
	SignWarning.Image = img
}
