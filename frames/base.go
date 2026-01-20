package frames

import "image"

type Definition struct {
	Name   string
	Middle image.Rectangle
	Image  image.Image
}
