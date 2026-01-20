package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed sign_warning_xlarge.png
var SignWarningXlargePng []byte

var SignWarningXlarge = &Definition{
	Name: "sign_warning_xlarge",
	Middle: image.Rect(36, 36, 156, 156),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(SignWarningXlargePng))
	if err != nil { panic(err) }
	SignWarningXlarge.Image = img
}
