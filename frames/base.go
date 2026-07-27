package frames

import "image"

// Definition represents a generated frame asset, including its image and middle geometry.
type Definition struct {
	// Name is the unique identifier for this frame design and size.
	Name   string
	// Middle defines the content area within the frame image that will be preserved or stretched.
	Middle image.Rectangle
	// Image contains the base pixel data for the frame border.
	Image  image.Image
}
