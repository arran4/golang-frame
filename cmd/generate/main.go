package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"sort"
)

type Generator func(scale int) (image.Image, image.Rectangle, string)

var generators = []Generator{
	// OS Like
	genWin31,
	genWin95,
	genMacClassic,
	genMacOSX,
	genMWM,
	genNeXT,
	genBeOS,
	genAmiga,

	// Windows
	genRetroWindow,
	genFutureWindow,
	genPaperWindow,
	genGlassWindow,

	// Actual / Material
	genWood,
	genGold,
	genMetal,
	genRidge,

	// Fancy / Pattern
	genFloral,
	genHearts,
	genStars,
	genCheckers,
	genDots,
	genWaves,
	genChains,
	genRainbow,

	// Ornate / Large
	genChinaPattern,
	genFantasyStone,
	genSciFiTech,

	// Signs
	genSignWarning,
	genSignStreet,
	genSignConstruction,
}

func main() {
	dstDir := "frames"
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		panic(err)
	}

	files, _ := filepath.Glob(filepath.Join(dstDir, "*"))
	for _, f := range files {
		os.Remove(f)
	}

	baseFile, _ := os.Create(filepath.Join(dstDir, "base.go"))
	fmt.Fprintln(baseFile, "package frames")
	fmt.Fprintln(baseFile, "")
	fmt.Fprintln(baseFile, "import \"image\"")
	fmt.Fprintln(baseFile, "")
	fmt.Fprintln(baseFile, "type Definition struct {")
	fmt.Fprintln(baseFile, "\tName   string")
	fmt.Fprintln(baseFile, "\tMiddle image.Rectangle")
	fmt.Fprintln(baseFile, "\tImage  image.Image")
	fmt.Fprintln(baseFile, "}")
	baseFile.Close()

	variants := []struct {
		Suffix string
		Scale  int
	}{
		{"", 1},
		{"_large", 2},
		{"_xlarge", 3},
	}

	var allExportedNames []string

	for _, gen := range generators {
		for _, variant := range variants {
			img, middle, baseName := gen(variant.Scale)
			name := baseName + variant.Suffix

			filename := name + ".png"
			f, _ := os.Create(filepath.Join(dstDir, filename))
			if err := png.Encode(f, img); err != nil {
				panic(err)
			}
			f.Close()

			exportedName := toExportedName(name)
			allExportedNames = append(allExportedNames, exportedName)
			goFilename := name + ".go"

			goFile, _ := os.Create(filepath.Join(dstDir, goFilename))
			fmt.Fprintln(goFile, "package frames")
			fmt.Fprintln(goFile, "")
			fmt.Fprintln(goFile, "import (")
			fmt.Fprintln(goFile, "\t_ \"embed\"")
			fmt.Fprintln(goFile, "\t\"image\"")
			fmt.Fprintln(goFile, "\t_ \"image/png\"")
			fmt.Fprintln(goFile, "\t\"bytes\"")
			fmt.Fprintln(goFile, ")")
			fmt.Fprintln(goFile, "")
			fmt.Fprintf(goFile, "//go:embed %s\n", filename)
			fmt.Fprintf(goFile, "var %sPng []byte\n\n", exportedName)
			fmt.Fprintf(goFile, "var %s = &Definition{\n", exportedName)
			fmt.Fprintf(goFile, "\tName: \"%s\",\n", name)
			fmt.Fprintf(goFile, "\tMiddle: image.Rect(%d, %d, %d, %d),\n", middle.Min.X, middle.Min.Y, middle.Max.X, middle.Max.Y)
			fmt.Fprintln(goFile, "}")
			fmt.Fprintln(goFile, "")
			fmt.Fprintln(goFile, "func init() {")
			fmt.Fprintf(goFile, "\timg, _, err := image.Decode(bytes.NewReader(%sPng))\n", exportedName)
			fmt.Fprintln(goFile, "\tif err != nil { panic(err) }")
			fmt.Fprintf(goFile, "\t%s.Image = img\n", exportedName)
			fmt.Fprintln(goFile, "}")
			goFile.Close()
		}
	}

	sort.Strings(allExportedNames)
	allFile, _ := os.Create(filepath.Join(dstDir, "all.go"))
	fmt.Fprintln(allFile, "package frames")
	fmt.Fprintln(allFile, "")
	fmt.Fprintln(allFile, "var All = []*Definition{")
	for _, en := range allExportedNames {
		fmt.Fprintf(allFile, "\t%s,\n", en)
	}
	fmt.Fprintln(allFile, "}")
	allFile.Close()
}

func toExportedName(s string) string {
	res := ""
	nextUpper := true
	for _, c := range s {
		if c == '_' {
			nextUpper = true
		} else {
			if nextUpper {
				if c >= 'a' && c <= 'z' {
					c = c - 32
				}
				nextUpper = false
			}
			res += string(c)
		}
	}
	return res
}

func solid(w, h int, c color.Color) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
	return img
}

func rect(img *image.RGBA, r image.Rectangle, c color.Color) {
	draw.Draw(img, r, &image.Uniform{c}, image.Point{}, draw.Src)
}

func rectHighlight(img *image.RGBA, r image.Rectangle, c color.RGBA) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if x < 0 || y < 0 || x >= img.Bounds().Dx() || y >= img.Bounds().Dy() {
				continue
			}
			old := img.RGBAAt(x, y)
			a := float64(c.A) / 255.0
			newR := uint8(float64(old.R)*(1-a) + float64(c.R)*a)
			newG := uint8(float64(old.G)*(1-a) + float64(c.G)*a)
			newB := uint8(float64(old.B)*(1-a) + float64(c.B)*a)
			img.SetRGBA(x, y, color.RGBA{newR, newG, newB, 255})
		}
	}
}

// --- GENERATORS ---

func genSignWarning(s int) (image.Image, image.Rectangle, string) {
	// Professional Industrial Hazard Container
	w, h := 64*s, 64*s
	img := solid(w, h, color.Transparent)
	bw := 12 * s
	yellow := color.RGBA{255, 204, 0, 255}
	black := color.RGBA{0, 0, 0, 255}
	darkPlate := color.RGBA{40, 40, 40, 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Border zone
			if x < bw || x >= w-bw || y < bw || y >= h-bw {
				// 45-degree Safety Stripes
				if (x+y)/(4*s)%2 == 0 {
					img.Set(x, y, yellow)
				} else {
					img.Set(x, y, black)
				}
			} else {
				// Inner plate
				img.Set(x, y, color.RGBA{255, 245, 200, 200})
			}
		}
	}
	// Riveted Corners
	rivetColor := color.RGBA{180, 180, 180, 255}
	drawRivet := func(cx, cy int) {
		rad := 2 * s
		for dy := -rad; dy <= rad; dy++ {
			for dx := -rad; dx <= rad; dx++ {
				if dx*dx+dy*dy <= rad*rad {
					img.Set(cx+dx, cy+dy, rivetColor)
				}
			}
		}
		img.Set(cx, cy, darkPlate)
	}
	drawRivet(bw/2, bw/2)
	drawRivet(w-bw/2, bw/2)
	drawRivet(bw/2, h-bw/2)
	drawRivet(w-bw/2, h-bw/2)

	return img, image.Rect(bw, bw, w-bw, h-bw), "sign_warning"
}

func genWood(s int) (image.Image, image.Rectangle, string) {
	w, h := 96*s, 96*s
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	c1, c2, c3 := color.RGBA{110, 60, 30, 255}, color.RGBA{80, 40, 20, 255}, color.RGBA{50, 25, 10, 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xf, yf := float64(x), float64(y)

			// Plank gaps every 24 pixels
			if (y/s)%24 < 2 {
				img.Set(x, y, color.RGBA{30, 15, 5, 255})
				continue
			}

			// Organic fibrous grain
			grain := math.Sin(yf/float64(s)+math.Sin(xf/float64(s*12))*5.0) * 0.5
			fineNoise := math.Sin(xf*10+yf*0.5) * 0.2

			// Knots
			knot := 0.0
			for _, k := range []image.Point{{w / 3, h / 4}, {2 * w / 3, 3 * h / 4}} {
				dx, dy := xf-float64(k.X), yf-float64(k.Y)
				d := math.Sqrt(dx*dx + dy*dy)
				knot += math.Exp(-d/float64(8*s)) * math.Sin(d/float64(s)) * 4.0
			}

			f := (grain + fineNoise + knot + 1.5) / 3.0
			if f < 0 {
				f = 0
			} else if f > 1 {
				f = 1
			}

			var c color.RGBA
			if f < 0.5 {
				t := f * 2
				c = color.RGBA{uint8(float64(c1.R)*(1-t) + float64(c2.R)*t), uint8(float64(c1.G)*(1-t) + float64(c2.G)*t), uint8(float64(c1.B)*(1-t) + float64(c2.B)*t), 255}
			} else {
				t := (f - 0.5) * 2
				c = color.RGBA{uint8(float64(c2.R)*(1-t) + float64(c3.R)*t), uint8(float64(c2.G)*(1-t) + float64(c3.G)*t), uint8(float64(c2.B)*(1-t) + float64(c3.B)*t), 255}
			}
			img.Set(x, y, c)
		}
	}
	bw := 16 * s
	for i := 0; i < bw; i++ {
		fade := uint8(40 - i*40/bw)
		rectHighlight(img, image.Rect(i, i, w-i, i+1), color.RGBA{255, 255, 255, fade})
		rectHighlight(img, image.Rect(0, i, i+1, h), color.RGBA{255, 255, 255, fade})
	}
	return img, image.Rect(bw, bw, w-bw, h-bw), "wood"
}

func genFloral(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{255, 250, 240, 255})
	pink := color.RGBA{255, 105, 180, 255}
	green := color.RGBA{34, 139, 34, 255}
	drawFlower := func(cx, cy, size int) {
		for i := 0; i < 5; i++ {
			angle := float64(i) * 2 * math.Pi / 5
			px := cx + int(math.Cos(angle)*float64(size))
			py := cy + int(math.Sin(angle)*float64(size))
			for dy := -size / 2; dy <= size/2; dy++ {
				for dx := -size / 2; dx <= size/2; dx++ {
					if dx*dx+dy*dy <= (size/2)*(size/2) {
						img.Set(px+dx, py+dy, pink)
					}
				}
			}
		}
		yellow := color.RGBA{255, 255, 0, 255}
		rad := size / 3
		for dy := -rad; dy <= rad; dy++ {
			for dx := -rad; dx <= rad; dx++ {
				if dx*dx+dy*dy <= rad*rad {
					img.Set(cx+dx, cy+dy, yellow)
				}
			}
		}
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x >= 16*s && x < w-16*s && y >= 16*s && y < h-16*s {
				continue
			}
			xf, yf := float64(x), float64(y)
			if math.Abs(math.Sin(xf/float64(8*s))*float64(4*s)+float64(h/2)-yf) < float64(s) {
				img.Set(x, y, green)
			}
		}
	}
	drawFlower(8*s, 8*s, 4*s)
	drawFlower(w-8*s, 8*s, 4*s)
	drawFlower(8*s, h-8*s, 4*s)
	drawFlower(w-8*s, h-8*s, 4*s)
	return img, image.Rect(16*s, 16*s, w-16*s, h-16*s), "floral"
}

func genStars(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			f := float64(y) / float64(h)
			img.Set(x, y, color.RGBA{uint8(5 + f*10), uint8(5 + f*10), uint8(20 + f*20), 255})
		}
	}
	for i := 0; i < 30*s; i++ {
		x, y := (i*137)%w, (i*149)%h
		c := color.RGBA{255, 255, 200, 255}
		img.Set(x, y, c)
		if s > 1 {
			rectHighlight(img, image.Rect(x-s, y, x+s+1, y+1), color.RGBA{c.R, c.G, c.B, 100})
			rectHighlight(img, image.Rect(x, y-s, x+1, y+s+1), color.RGBA{c.R, c.G, c.B, 100})
		}
	}
	return img, image.Rect(8*s, 8*s, w-8*s, h-8*s), "stars"
}

func genGold(s int) (image.Image, image.Rectangle, string) {
	w, h := 96*s, 96*s
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bw := 16 * s

	baseColor := color.RGBA{218, 165, 32, 255} // Metallic Gold

	// Profile function: returns height (0-1) at normalized distance t (0-1)
	getProfile := func(t float64) float64 {
		if t < 0.0 {
			return 0.0
		}
		if t > 1.0 {
			return 0.0
		}
		// Classic Ogee / Scoop profile
		// 0.0 - 0.15: Outer Bead
		if t < 0.15 {
			tt := t / 0.15
			return math.Sqrt(1-(tt-1)*(tt-1)) * 0.8 // Quarter circle
		}
		// 0.15 - 0.20: Step down
		if t < 0.20 {
			tt := (t - 0.15) / 0.05
			return 0.8 - tt*0.2
		}
		// 0.20 - 0.70: The Scoop (Concave)
		if t < 0.70 {
			tt := (t - 0.20) / 0.50
			// concave curve
			return 0.6 - math.Sin(tt*math.Pi)*0.3
		}
		// 0.70 - 0.85: Inner Bead (Convex)
		if t < 0.85 {
			tt := (t - 0.70) / 0.15
			return 0.6 + math.Sin(tt*math.Pi)*0.4
		}
		// 0.85 - 1.00: Step to picture
		tt := (t - 0.85) / 0.15
		return 0.6 * (1 - tt)
	}

	lx, ly, lz := -1.0, -1.0, 0.5
	ln := math.Sqrt(lx*lx + ly*ly + lz*lz)
	lx, ly, lz = lx/ln, ly/ln, lz/ln

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Calculate distance to edge and gradient of distance
			d := x
			gx, gy := 1.0, 0.0
			if w-1-x < d {
				d = w - 1 - x
				gx, gy = -1.0, 0.0
			}
			if y < d {
				d = y
				gx, gy = 0.0, 1.0
			}
			if h-1-y < d {
				d = h - 1 - y
				gx, gy = 0.0, -1.0
			}

			if d >= bw {
				continue
			}

			t := float64(d) / float64(bw)

			// Numerical derivative
			z1 := getProfile(t)
			z2 := getProfile(t + 0.01)
			slope := (z2 - z1) / 0.01

			// Normal calc
			heightScale := float64(bw) * 0.5
			realSlope := slope * heightScale / float64(bw) // dz/dt * Zscale / (Dscale)

			nx := -realSlope * gx
			ny := -realSlope * gy
			nz := 1.0

			// Texture / Bump map
			noiseScale := 0.2
			n1 := math.Sin(float64(x)*0.4) * math.Cos(float64(y)*0.4)
			n2 := math.Cos(float64(x)*0.7 + float64(y)*0.7)
			nx += n1 * noiseScale
			ny += n2 * noiseScale

			nn := math.Sqrt(nx*nx + ny*ny + nz*nz)
			nx, ny, nz = nx/nn, ny/nn, nz/nn

			// Diffuse
			dot := nx*lx + ny*ly + nz*lz
			if dot < 0 {
				dot = 0
			}

			// Specular
			spec := 0.0
			refZ := 2*dot*nz - lz
			if refZ > 0 {
				spec = math.Pow(refZ, 20) // shininess
			}

			// Composite color
			// Ambient
			r := float64(baseColor.R) * 0.4
			g := float64(baseColor.G) * 0.4
			b := float64(baseColor.B) * 0.4

			// Diffuse
			r += float64(baseColor.R) * 0.6 * dot
			g += float64(baseColor.G) * 0.6 * dot
			b += float64(baseColor.B) * 0.6 * dot

			// Specular (white)
			r += 255 * spec * 0.4
			g += 255 * spec * 0.4
			b += 255 * spec * 0.4

			if r > 255 {
				r = 255
			}
			if g > 255 {
				g = 255
			}
			if b > 255 {
				b = 255
			}

			img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}
	return img, image.Rect(bw, bw, w-bw, h-bw), "gold"
}

func genSignConstruction(s int) (image.Image, image.Rectangle, string) {
	w, h := 96*s, 96*s
	orange := color.RGBA{255, 120, 0, 255}
	img := solid(w, h, orange)
	bw := 16 * s
	black := color.RGBA{0, 0, 0, 255}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x < bw || x >= w-bw || y < bw || y >= h-bw {
				if (x+y)/(8*s)%2 == 0 {
					img.Set(x, y, black)
				}
			}
		}
	}
	rectHighlight(img, image.Rect(bw-s, bw-s, w-bw+s, bw), black)
	return img, image.Rect(bw, bw, w-bw, h-bw), "sign_construction"
}

func genChinaPattern(s int) (image.Image, image.Rectangle, string) {
	w, h := 128*s, 128*s
	img := solid(w, h, color.White)
	cobalt := color.RGBA{0, 71, 171, 255}
	drawQuarter := func(offsetX, offsetY int, flipX, flipY bool) {
		for y := 0; y < 48*s; y++ {
			for x := 0; x < 48*s; x++ {
				xf, yf := float64(x)/float64(s), float64(y)/float64(s)
				val := math.Sin(xf*0.3)*math.Cos(yf*0.3)*10.0 + math.Sin(xf*yf*0.05)*5.0
				cx, cy := x, y
				if flipX {
					cx = 48*s - 1 - x
				}
				if flipY {
					cy = 48*s - 1 - y
				}
				if int(val)%7 == 0 && (x+y)%2 == 0 {
					img.Set(offsetX+cx, offsetY+cy, cobalt)
				}
				if x < 4*s || y < 4*s {
					img.Set(offsetX+cx, offsetY+cy, cobalt)
				}
				if x > 44*s && x < 46*s || y > 44*s && y < 46*s {
					img.Set(offsetX+cx, offsetY+cy, cobalt)
				}
			}
		}
	}
	drawQuarter(0, 0, false, false)
	drawQuarter(w-48*s, 0, true, false)
	drawQuarter(0, h-48*s, false, true)
	drawQuarter(w-48*s, h-48*s, true, true)
	for i := 48 * s; i < w-48*s; i++ {
		if (i/s)%8 < 4 {
			for j := 0; j < 4*s; j++ {
				img.Set(i, j, cobalt)
				img.Set(i, h-1-j, cobalt)
				img.Set(j, i, cobalt)
				img.Set(w-1-j, i, cobalt)
			}
		}
	}
	return img, image.Rect(48*s, 48*s, w-48*s, h-48*s), "china_pattern"
}

func genFutureWindow(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{10, 10, 25, 240})
	cyan := color.RGBA{0, 255, 255, 255}
	rectHighlight(img, image.Rect(0, 0, w, s), cyan)
	rectHighlight(img, image.Rect(w-s, 0, w, h), cyan)
	rectHighlight(img, image.Rect(w-8*s, 2*s, w-2*s, 6*s), cyan)
	return img, image.Rect(12*s, 24*s, w-12*s, h-12*s), "window_future"
}

func genCheckers(s int) (image.Image, image.Rectangle, string) {
	sq := 8 * s
	w, h := sq*8, sq*8
	img := solid(w, h, color.White)
	black := color.RGBA{40, 40, 40, 255}
	for y := 0; y < h; y += sq {
		for x := 0; x < w; x += sq {
			if (x/sq+y/sq)%2 == 1 {
				rect(img, image.Rect(x, y, x+sq, y+sq), black)
			}
		}
	}
	return img, image.Rect(sq*2, sq*2, sq*4, sq*4), "checkers"
}

func genDots(s int) (image.Image, image.Rectangle, string) {
	grid := 16 * s
	w, h := grid*4, grid*4
	img := solid(w, h, color.White)
	blue := color.RGBA{0, 100, 255, 255}
	for y := grid / 2; y < h; y += grid {
		for x := grid / 2; x < w; x += grid {
			for dy := -2 * s; dy <= 2*s; dy++ {
				for dx := -2 * s; dx <= 2*s; dx++ {
					if dx*dx+dy*dy <= 4*s*s {
						img.Set(x+dx, y+dy, blue)
					}
				}
			}
		}
	}
	return img, image.Rect(grid, grid, grid*2, grid*2), "dots"
}

func genWin31(s int) (image.Image, image.Rectangle, string) {
	w, h := 32*s, 32*s
	img := solid(w, h, color.RGBA{192, 192, 192, 255})
	rect(img, image.Rect(0, 0, w, s), color.White)
	rect(img, image.Rect(w-s, 0, w, h), color.RGBA{80, 80, 80, 255})
	return img, image.Rect(4*s, 4*s, w-4*s, h-4*s), "win31_like"
}

func genWin95(s int) (image.Image, image.Rectangle, string) {
	w, h := 32*s, 32*s
	img := solid(w, h, color.RGBA{192, 192, 192, 255})
	rect(img, image.Rect(0, 0, w, s), color.White)
	rect(img, image.Rect(w-s, 0, w, h), color.Black)
	return img, image.Rect(4*s, 4*s, w-4*s, h-4*s), "win95_like"
}

func genMacClassic(s int) (image.Image, image.Rectangle, string) {
	w, h := 48*s, 48*s
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// Outline
	rect(img, image.Rect(0, 0, w, s), black)
	rect(img, image.Rect(0, h-s, w, h), black)
	rect(img, image.Rect(0, 0, s, h), black)
	rect(img, image.Rect(w-s, 0, w, h), black)

	// Title bar
	titleBarH := 21 * s
	rect(img, image.Rect(0, titleBarH, w, titleBarH+s), black)

	// Stripes
	for y := 2 * s; y < titleBarH; y += 2 * s {
		rect(img, image.Rect(s, y, w-s, y+s), black)
	}

	// Close box
	cbSize := 11 * s
	cbX := 6 * s
	cbY := 5 * s

	rect(img, image.Rect(cbX, cbY, cbX+cbSize, cbY+cbSize), white)
	rect(img, image.Rect(cbX, cbY, cbX+cbSize, cbY+s), black)
	rect(img, image.Rect(cbX, cbY+cbSize-s, cbX+cbSize, cbY+cbSize), black)
	rect(img, image.Rect(cbX, cbY, cbX+s, cbY+cbSize), black)
	rect(img, image.Rect(cbX+cbSize-s, cbY, cbX+cbSize, cbY+cbSize), black)

	return img, image.Rect(20*s, 22*s, w-4*s, h-4*s), "mac_classic_like"
}

func genMacOSX(s int) (image.Image, image.Rectangle, string) {
	width := 64 * s
	height := 64 * s
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	titleBarStart := color.RGBA{235, 235, 235, 255}
	titleBarEnd := color.RGBA{210, 210, 210, 255}
	borderColor := color.RGBA{180, 180, 180, 255}

	middle := image.Rect(54*s, 24*s, width-4*s, height-4*s)
	cornerRadius := 8 * s
	smallRad := 4 * s

	// Draw Title Bar Background
	for y := 0; y < 24*s; y++ {
		f := float64(y) / float64(24*s)
		c := interpolate(titleBarStart, titleBarEnd, f)
		rect(img, image.Rect(0, y, width, y+1), c)
	}

	// Draw Window Body Background (White)
	rect(img, image.Rect(0, 24*s, width, height), color.White)

	// Masking out the corners
	for y := 0; y < cornerRadius; y++ {
		for x := 0; x < cornerRadius; x++ {
			dx := cornerRadius - 1 - x
			dy := cornerRadius - 1 - y
			if dx*dx+dy*dy >= cornerRadius*cornerRadius {
				img.Set(x, y, color.Transparent)
			}
		}
	}
	for y := 0; y < cornerRadius; y++ {
		for x := 0; x < cornerRadius; x++ {
			dx := x
			dy := cornerRadius - 1 - y
			if dx*dx+dy*dy >= cornerRadius*cornerRadius {
				img.Set(width-cornerRadius+x, y, color.Transparent)
			}
		}
	}
	for y := 0; y < smallRad; y++ {
		for x := 0; x < smallRad; x++ {
			dx := smallRad - 1 - x
			dy := y
			if dx*dx+dy*dy >= smallRad*smallRad {
				img.Set(x, height-smallRad+y, color.Transparent)
			}
		}
	}
	for y := 0; y < smallRad; y++ {
		for x := 0; x < smallRad; x++ {
			dx := x
			dy := y
			if dx*dx+dy*dy >= smallRad*smallRad {
				img.Set(width-smallRad+x, height-smallRad+y, color.Transparent)
			}
		}
	}

	addBorder(img, borderColor)

	btnY := 12 * s
	btnRad := 5 * s
	gap := 8 * s
	startX := 10 * s
	red := color.RGBA{255, 95, 87, 255}
	yellow := color.RGBA{255, 189, 46, 255}
	green := color.RGBA{40, 201, 64, 255}

	drawCircle(img, startX, btnY, btnRad, red)
	drawCircle(img, startX+gap+2*btnRad, btnY, btnRad, yellow)
	drawCircle(img, startX+2*gap+4*btnRad, btnY, btnRad, green)

	return img, middle, "macosx_like"
}

func genMWM(s int) (image.Image, image.Rectangle, string) {
	w, h := 48*s, 48*s
	img := solid(w, h, color.RGBA{180, 180, 180, 255})
	rect(img, image.Rect(0, 0, w, 4*s), color.White)
	return img, image.Rect(6*s, 6*s, w-6*s, h-6*s), "mwm_like"
}

func genNeXT(s int) (image.Image, image.Rectangle, string) {
	w, h := 32*s, 32*s
	img := solid(w, h, color.Black)
	rect(img, image.Rect(s, s, w-s, h-s), color.RGBA{150, 150, 150, 255})
	return img, image.Rect(6*s, 6*s, w-6*s, h-6*s), "next_like"
}

func genBeOS(s int) (image.Image, image.Rectangle, string) {
	w, h := 32*s, 32*s
	img := solid(w, h, color.RGBA{255, 204, 0, 255})
	return img, image.Rect(4*s, 4*s, w-4*s, h-4*s), "beos_like"
}

func genAmiga(s int) (image.Image, image.Rectangle, string) {
	w, h := 32*s, 32*s
	img := solid(w, h, color.RGBA{0, 80, 160, 255})
	return img, image.Rect(6*s, 6*s, w-6*s, h-6*s), "amiga_like"
}

func genRetroWindow(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{192, 192, 192, 255})
	return img, image.Rect(8*s, 20*s, w-8*s, h-8*s), "window_retro"
}

func genPaperWindow(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.Transparent)
	rect(img, image.Rect(0, 0, w-4*s, h-4*s), color.RGBA{255, 253, 240, 255})
	return img, image.Rect(8*s, 16*s, w-12*s, h-12*s), "window_paper"
}

func genGlassWindow(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{255, 255, 255, 80})
	return img, image.Rect(8*s, 18*s, w-8*s, h-8*s), "window_glass"
}

func genMetal(s int) (image.Image, image.Rectangle, string) {
	w, h := 48*s, 48*s
	img := solid(w, h, color.RGBA{180, 180, 190, 255})
	return img, image.Rect(8*s, 8*s, w-8*s, h-8*s), "metal"
}

func genRidge(s int) (image.Image, image.Rectangle, string) {
	w, h := 48*s, 48*s
	img := solid(w, h, color.RGBA{200, 200, 200, 255})
	return img, image.Rect(8*s, 8*s, w-8*s, h-8*s), "ridge"
}

func genHearts(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{255, 240, 245, 255})
	red := color.RGBA{220, 20, 60, 255}
	drawHeart := func(cx, cy, size int) {
		for y := -size; y <= size; y++ {
			for x := -size; x <= size; x++ {
				xf, yf := float64(x)/float64(size), float64(y)/float64(size)
				if math.Pow(xf*xf+yf*yf-1, 3)-xf*xf*yf*yf*yf <= 0 {
					img.Set(cx+x, cy-y, red)
				}
			}
		}
	}
	drawHeart(16*s, 16*s, 6*s)
	return img, image.Rect(16*s, 16*s, w-16*s, h-16*s), "hearts"
}

func genWaves(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{0, 105, 148, 255})
	return img, image.Rect(8*s, 8*s, w-8*s, h-8*s), "waves"
}

func genChains(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.White)
	return img, image.Rect(8*s, 8*s, w-8*s, h-8*s), "chains"
}

func genRainbow(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.White)
	return img, image.Rect(8*s, 16*s, w-8*s, h-8*s), "rainbow"
}

func genFantasyStone(s int) (image.Image, image.Rectangle, string) {
	w, h := 96*s, 96*s
	img := solid(w, h, color.RGBA{80, 80, 80, 255})
	return img, image.Rect(20*s, 20*s, w-20*s, h-20*s), "fantasy_stone"
}

func genSciFiTech(s int) (image.Image, image.Rectangle, string) {
	w, h := 96*s, 96*s
	img := solid(w, h, color.Black)
	rectHighlight(img, image.Rect(0, 0, w, 4*s), color.RGBA{0, 200, 255, 255})
	return img, image.Rect(16*s, 16*s, w-16*s, h-16*s), "scifi_tech"
}

func genSignStreet(s int) (image.Image, image.Rectangle, string) {
	w, h := 64*s, 64*s
	img := solid(w, h, color.RGBA{0, 100, 0, 255})
	return img, image.Rect(8*s, 8*s, w-8*s, h-8*s), "sign_street"
}

// Helpers

func interpolate(c1, c2 color.RGBA, f float64) color.RGBA {
	return color.RGBA{
		uint8(float64(c1.R)*(1-f) + float64(c2.R)*f),
		uint8(float64(c1.G)*(1-f) + float64(c2.G)*f),
		uint8(float64(c1.B)*(1-f) + float64(c2.B)*f),
		uint8(float64(c1.A)*(1-f) + float64(c2.A)*f),
	}
}

func drawCircle(img *image.RGBA, cx, cy, r int, c color.RGBA) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				if cx+x >= 0 && cx+x < img.Bounds().Dx() && cy+y >= 0 && cy+y < img.Bounds().Dy() {
					img.Set(cx+x, cy+y, c)
				}
			}
		}
	}
}

func addBorder(img *image.RGBA, c color.RGBA) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	var boundary []image.Point

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			} // Transparent

			isBoundary := false
			if x == 0 || x == w-1 || y == 0 || y == h-1 {
				isBoundary = true
			} else {
				// Check 4 neighbors
				if _, _, _, na := img.At(x-1, y).RGBA(); na == 0 {
					isBoundary = true
				} else if _, _, _, na := img.At(x+1, y).RGBA(); na == 0 {
					isBoundary = true
				} else if _, _, _, na := img.At(x, y-1).RGBA(); na == 0 {
					isBoundary = true
				} else if _, _, _, na := img.At(x, y+1).RGBA(); na == 0 {
					isBoundary = true
				}
			}

			if isBoundary {
				boundary = append(boundary, image.Point{x, y})
			}
		}
	}

	for _, p := range boundary {
		img.Set(p.X, p.Y, c)
	}
}
