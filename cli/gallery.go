package cli

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"text/template"

	frame "github.com/arran4/golang-frame"
	"github.com/arran4/golang-frame/frames"
)

type FrameData struct {
	Name         string
	ExportedName string
	IsCheckers   bool
}

// Gallery is a subcommand `frames gallery` Generates gallery images and readme
func Gallery() error {
	dstDir := "images"
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	var frameDatas []FrameData

	for _, def := range frames.All {
		wLow, _, hLow, _ := frame.MeasureFrame(def.Image, def.Middle, 180, 100)
		w, h := wLow+60, hLow+60
		dst := image.NewRGBA(image.Rect(0, 0, w, h))
		// Light transparent grey background to show transparency
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if (x/8+y/8)%2 == 0 {
					dst.Set(x, y, color.RGBA{240, 240, 240, 255})
				} else {
					dst.Set(x, y, color.RGBA{255, 255, 255, 255})
				}
			}
		}

		// Create frame
		rect := image.Rect(30, 30, 30+wLow, 30+hLow)
		fr := frame.NewFrame(rect, def.Image, def.Middle)
		draw.Draw(dst, rect, fr, rect.Min, draw.Over)

		filename := fmt.Sprintf("gallery_%s.png", def.Name)
		f, err := os.Create(filepath.Join(dstDir, filename))
		if err != nil {
			return err
		}
		if err := png.Encode(f, dst); err != nil {
			return err
		}
		f.Close()

		exportedName := toExportedName(def.Name)

		// If checkers, generate additional aspect ratio / size examples
		if def.Name == "checkers" {
			for idx, wh := range [][2]int{{300, 150}, {150, 300}, {183, 201}, {230, 230}} {
				// Measure it properly!
				wLow, _, hLow, _ := frame.MeasureFrame(def.Image, def.Middle, wh[0], wh[1])
				w, h := wLow, hLow
				dstExtra := image.NewRGBA(image.Rect(0, 0, w+60, h+60))
				for y := 0; y < h+60; y++ {
					for x := 0; x < w+60; x++ {
						if (x/8+y/8)%2 == 0 {
							dstExtra.Set(x, y, color.RGBA{240, 240, 240, 255})
						} else {
							dstExtra.Set(x, y, color.RGBA{255, 255, 255, 255})
						}
					}
				}
				rectExtra := image.Rect(30, 30, 30+w, 30+h)
				frExtra := frame.NewFrame(rectExtra, def.Image, def.Middle)
				draw.Draw(dstExtra, rectExtra, frExtra, rectExtra.Min, draw.Over)

				extraFilename := fmt.Sprintf("gallery_%s_extra_%d.png", def.Name, idx)
				fe, err := os.Create(filepath.Join(dstDir, extraFilename))
				if err != nil {
					return err
				}
				if err := png.Encode(fe, dstExtra); err != nil {
					fe.Close()
					return err
				}
				fe.Close()
			}
		}

		frameDatas = append(frameDatas, FrameData{
			Name:         def.Name,
			ExportedName: exportedName,
			IsCheckers:   def.Name == "checkers",
		})
	}

	tmpl, err := template.ParseFiles("readme.md.tmpl")
	if err != nil {
		return err
	}

	readmeFile, err := os.Create("readme.md")
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	data := struct {
		Frames []FrameData
	}{
		Frames: frameDatas,
	}

	if err := tmpl.Execute(readmeFile, data); err != nil {
		return err
	}

	fmt.Println("Successfully generated images/ and readme.md")
	return nil
}
