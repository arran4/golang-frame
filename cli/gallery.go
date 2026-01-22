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
}

// Gallery is a subcommand `frames gallery` Generates gallery images and readme
func Gallery() error {
	dstDir := "images"
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	var frameDatas []FrameData

	for _, def := range frames.All {
		// Create a sample image for this frame
		w, h := 240, 160
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
		rect := image.Rect(30, 30, w-30, h-30)
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
		frameDatas = append(frameDatas, FrameData{
			Name:         def.Name,
			ExportedName: exportedName,
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
