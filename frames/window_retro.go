package frames

import (
	_ "embed"
	"image"
	_ "image/png"
	"bytes"
)

//go:embed window_retro.png
var WindowRetroPng []byte

var WindowRetro = &Definition{
	Name: "window_retro",
	Middle: image.Rect(22, 25, 38, 90),
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(WindowRetroPng))
	if err != nil { panic(err) }
	WindowRetro.Image = img
}
