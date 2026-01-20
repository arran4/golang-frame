package cli

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// GenWood is a subcommand `frames wood`
func GenWood() {
	generateWoodFile("frames/wood.png", 96, 16)
	generateWoodFile("frames/wood_large.png", 192, 32)
	generateWoodFile("frames/wood_xlarge.png", 288, 48)
}

func generateWoodFile(filename string, size int, border int) {
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	// 1. Generate Background (Inner part)
	// We fill the whole image first.
	// Background wood: darker, vertical grain.
	fillWood(img, image.Rect(0,0,size,size), true, 0.5)

	// 2. Generate Frame Borders
	// We iterate over the border pixels.
	w, h := size, size

	// Frame appearance parameters
	frameBaseColor := color.RGBA{139, 69, 19, 255} // SaddleBrown

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Check if we are in the border region
			isBorder := false
			if x < border || x >= w-border || y < border || y >= h-border {
				isBorder = true
			}

			if !isBorder {
				continue // Keep background
			}

			// Mitered corners logic
			var verticalGrain bool

			// Diagonals for miter: y=x and y=h-1-x
            // Top: y < x && y < h-1-x
            // Bottom: y > x && y > h-1-x
            // Left: x < y && x < h-1-y
            // Right: x > y && x > h-1-y

			// To handle equality (on the line), we favor top/bottom or left/right

			if y < x && y < h-1-x {
				// Top
				verticalGrain = false // Horizontal grain
			} else if y > x && y > h-1-x {
				// Bottom
				verticalGrain = false
			} else {
				// Left or Right (or on diagonals)
				verticalGrain = true
			}

            // Generate wood color for this pixel
            c := woodColor(x, y, verticalGrain, frameBaseColor)

            // Apply lighting for 3D effect
            // Top and Left get lighter (highlight)
            // Bottom and Right get darker (shadow)

            if y < x && y < h-1-x {
                 c = lighten(c, 0.2) // Top
            } else if x < y && x < h-1-y {
                 c = lighten(c, 0.2) // Left
            } else {
                 c = darken(c, 0.2) // Bottom, Right
            }

            img.Set(x, y, c)
		}
	}

	// Add distinct dark lines at miter joints
	drawMiterLines(img, size, border)

    // Add inner and outer frame borders (rectangles) to define the edge
    drawRect(img, image.Rect(0,0,size,size), color.RGBA{60,30,0,255}) // Outer
    drawRect(img, image.Rect(border, border, size-border, size-border), color.RGBA{60,30,0,255}) // Inner

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func woodColor(x, y int, vertical bool, base color.RGBA) color.RGBA {
    // Coordinate transformation for grain
    nx, ny := float64(x), float64(y)

    var val float64
    if vertical {
       // Grain vertical. Stretch along Y.
       val = turbulence(nx * 0.2, ny * 0.02)
    } else {
       // Grain horizontal. Stretch along X.
       val = turbulence(nx * 0.02, ny * 0.2)
    }

    // Sine wave for rings/plank feel
    var across float64
    if vertical {
        across = nx
    } else {
        across = ny
    }

    // freq 0.3 means approx 3 pixels per wave? no. 2pi/0.3 = 20 pixels.
    grain := math.Sin(across * 0.3 + val * 4.0)

    // Normalize -1..1 to 0..1
    grain = (grain + 1.0) / 2.0

    // Perturb base color
    r := float64(base.R) * (0.6 + 0.4*grain)
    g := float64(base.G) * (0.6 + 0.4*grain)
    b := float64(base.B) * (0.6 + 0.4*grain)

    return color.RGBA{clamp(r), clamp(g), clamp(b), 255}
}

func turbulence(x, y float64) float64 {
    val := 0.0
    val += noise(x, y) / 2.0
    val += noise(x*2, y*2) / 4.0
    val += noise(x*4, y*4) / 8.0
    return val
}

// Simple pseudo-random noise
func noise(x, y float64) float64 {
    // Integer part
    ix := int(math.Floor(x))
    iy := int(math.Floor(y))
    // Fractional part
    fx := x - float64(ix)
    fy := y - float64(iy)

    // Smooth step
    sx := fx * fx * (3 - 2*fx)
    sy := fy * fy * (3 - 2*fy)

    // Random values at corners
    n00 := randVal(ix, iy)
    n10 := randVal(ix+1, iy)
    n01 := randVal(ix, iy+1)
    n11 := randVal(ix+1, iy+1)

    // Interpolate
    nx0 := lerp(n00, n10, sx)
    nx1 := lerp(n01, n11, sx)

    return lerp(nx0, nx1, sy)
}

func randVal(x, y int) float64 {
    // Deterministic random based on coordinates
    seed := x * 57 + y * 131
    seed = (seed << 13) ^ seed
    res := (1.0 - float64((seed * (seed * seed * 15731 + 789221) + 1376312589) & 0x7fffffff) / 1073741824.0)
    return res
}

func lerp(a, b, t float64) float64 {
    return a + t*(b-a)
}

func clamp(v float64) uint8 {
    if v < 0 { return 0 }
    if v > 255 { return 255 }
    return uint8(v)
}

func lighten(c color.RGBA, amount float64) color.RGBA {
     r := float64(c.R) * (1.0 + amount)
     g := float64(c.G) * (1.0 + amount)
     b := float64(c.B) * (1.0 + amount)
     return color.RGBA{clamp(r), clamp(g), clamp(b), c.A}
}

func darken(c color.RGBA, amount float64) color.RGBA {
     r := float64(c.R) * (1.0 - amount)
     g := float64(c.G) * (1.0 - amount)
     b := float64(c.B) * (1.0 - amount)
     return color.RGBA{clamp(r), clamp(g), clamp(b), c.A}
}

func fillWood(img *image.RGBA, r image.Rectangle, vertical bool, seedOffset float64) {
    base := color.RGBA{100, 60, 20, 255} // Darker for background
     for y := r.Min.Y; y < r.Max.Y; y++ {
        for x := r.Min.X; x < r.Max.X; x++ {
             c := woodColor(x + int(seedOffset*1000), y, vertical, base)
             img.Set(x, y, c)
        }
     }
}

func drawMiterLines(img *image.RGBA, size int, border int) {
    // Draw diagonal lines at corners
    col := color.RGBA{40, 20, 0, 100} // Semi-transparent dark

    // Only need to draw within the border area
    // Top-Left corner
    for i := 0; i < border; i++ {
        img.Set(i, i, blend(img.RGBAAt(i, i), col))
    }
    // Top-Right corner
    for i := 0; i < border; i++ {
        x, y := size-1-i, i
        img.Set(x, y, blend(img.RGBAAt(x, y), col))
    }
    // Bottom-Left corner
    for i := 0; i < border; i++ {
        x, y := i, size-1-i
        img.Set(x, y, blend(img.RGBAAt(x, y), col))
    }
    // Bottom-Right corner
    for i := 0; i < border; i++ {
        x, y := size-1-i, size-1-i
        img.Set(x, y, blend(img.RGBAAt(x, y), col))
    }
}

func blend(c1, c2 color.RGBA) color.RGBA {
    // Alpha blending
    a := float64(c2.A) / 255.0
    r := float64(c1.R)*(1-a) + float64(c2.R)*a
    g := float64(c1.G)*(1-a) + float64(c2.G)*a
    b := float64(c1.B)*(1-a) + float64(c2.B)*a
    return color.RGBA{clamp(r), clamp(g), clamp(b), 255}
}

func drawRect(img *image.RGBA, r image.Rectangle, col color.RGBA) {
    // Top
    for x := r.Min.X; x < r.Max.X; x++ {
        img.Set(x, r.Min.Y, col)
        img.Set(x, r.Max.Y-1, col)
    }
    // Left
    for y := r.Min.Y; y < r.Max.Y; y++ {
        img.Set(r.Min.X, y, col)
        img.Set(r.Max.X-1, y, col)
    }
}
