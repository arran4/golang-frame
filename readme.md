[//]: # (This file is generated. Do not edit manually.)
# Frame

A simple golang library that creates an image.Image compatible object (ideally used with draw.Draw) that can be used for drawing a frame around something else. Think like a window, button, or anything else. 

It works by splitting an image into 9 parts:

![](doc/9%20parts.png)

Parts, 1, 3, 7 and 9 are drawn exactly as is. Then the remaining parts
are repeated or stretched to match the content as required.

Other variants might be created such as ones which are more procedurally generated or allow more interesting
views.

# Contributions

This is a simple library I am reusing else where. PRs reviewed and probably accepted. But definitely appreciated.

# Doc

There is basic go doc which you can check out here:
https://pkg.go.dev/github.com/arran4/golang-frame

The intended way of using this is:
```go
	fr := NewBasicFrame(targetArea)
	dst := i.SubImage(targetArea).(draw.Image)
	draw.Draw(dst, dst.Bounds(), fr, dst.Bounds().Min, draw.Src)
```

# Usage

## Sample 2: Drawing borders

Simplest possible use case; take an image and expand it to fit the 
desired size.

```go
	base, err := png.Decode(bytes.NewReader(baseImageData))
	if err != nil {
		log.Panicf("Error with loading base file: %s", err)
	}
	i := image.NewRGBA(image.Rect(0, 0, 600, 600))
	dst := i.SubImage(image.Rect(100, 100, 400, 400)).(draw.Image)
	fr := frame.NewFrame(dst.Bounds(), base, image.Rect(48,48,55, 66))
	draw.Draw(dst, dst.Bounds(), fr, dst.Bounds().Min, draw.Src)
	SaveFile(i)
```

Which will produce:

![](images/sample2.png)

From:

![](example/sample2/window.png)

## Sample 3: Section 5 image

This creates a better window implementation than sample 2. But also it shows that the way the borders
are drawn can be changed. As you can almost see in sample 2 the borders aren't drawn well as it's simply
repeating the contents. You can also use a (simple) stretch version instead of the repeating version. 
This is done with the `BorderMode` options see:
```go
	fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66), frame.Repeating)
	fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66), frame.Stretched)
```
See the sample for a more detailed look at the code. However, the difference this creates is as follows:

![](images/sample3.png)

From: 

![](example/sample3/window.png)

## Sample 4

In sample 4 we replace the contents of the window with our own rather than use section 5 of the image.

There are currently 3 variants of this:
* Section5Zeroed - Match section 5 starting position with co-ordinates 0, 0
* Zerod - Match the whole frame's starting position with the co-ordinates 0, 0
* PassThrough - Pass in the parent windows position

```go
fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66), &frame.Section5{Image: s5i}, frame.Section5Zeroed)
fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66), &frame.Section5{Image: s5i}, frame.Zerod)
fr := frame.NewFrame(frdst.Bounds(), base.(SubImagable).SubImage(s2), image.Rect(14, 48, 88, 66), &frame.Section5{Image: s5i}, frame.PassThrough) 
```

Which draws:

![](images/sample4.png)

From the frame:

![](example/sample4/window.png)

And with the section 5 image: 

![](example/sample4/person.png)

Please note, currently there is no support / consideration for an image with a none 0,0 Rectangle.Min
position. This might change so ensure your code will handle this. 

### Sample 4: Simple static image

```go

func NewBasicFrame(r image.Rectangle) *Frame {
	middle := image.Rect(0, 0, 1, 1)
	base := image.NewRGBA(image.Rect(-2, -2, 2, 2))
	b := base.Bounds()
	for y, r := range [][]color.RGBA{
		{colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray},
		{colornames.Lightgray, colornames.Darkgrey, colornames.Darkgrey, colornames.Darkgrey, colornames.Lightgray},
		{colornames.Lightgray, colornames.Darkgrey, colornames.White, colornames.Darkgrey, colornames.Lightgray},
		{colornames.Lightgray, colornames.Darkgrey, colornames.Darkgrey, colornames.Darkgrey, colornames.Lightgray},
		{colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray, colornames.Lightgray},
	} {
		for x, c := range r {
			base.Set(b.Min.X + x, b.Min.Y + y, c)
		}
	}
	return NewFrame(r, base, middle)
}

func main() {
    i := image.NewRGBA(image.Rect(0, 0, 150, 100))
    targetArea := image.Rect(10, 10, 100, 30)
    fr := NewBasicFrame(targetArea)
    dst := i.SubImage(targetArea).(draw.Image)
    draw.Draw(dst, dst.Bounds(), fr, dst.Bounds().Min, draw.Src)
}

```

Noting really amazing here but no need to have files, if you just want to draw a simple border you can 
do it this way, you might be able to wrap `image.NewUniform(color goes here).Bounds()` with a more restricted version. 

## Sample 5:

Section5 overlaying the previous image. Section5 now allows you to either replace, or overlay (with alpha and all) the 
base image's section 5.

Such as:
```go
	fr := frame.NewFrame(frdst.Bounds(), fibase.(SubImagable), image.Rect(11, 11, 111, 97), &frame.Section5{Image: s5i, Replace: false}, frame.Section5Zeroed, frame.Stretched)
```

![](images/sample5.png)

From:

![](example/sample5/frame.png)

and the source image generated by:

```go
	s5i := image.NewRGBA(image.Rect(0, 0, 50, 50))
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			if x/10%2 == 0 && x/10 == y/10 {
				s5i.SetRGBA(x, y, color.RGBA{0, 0, 127, 127})
			}
		}
	}
```

## Additional Helper functions

### `Draw.Over()`

`Draw.Over()` is an alpha function for Section5 overlays. (It's also the default if you don't specify one) 

Usage:
```go
import "github.com/arran4/golang-frame/draw"

frame.Section5{Image: s5i, Replace: false, AlphaMode: draw.Over}

```

# Included Frames

There is a collection of included frames in the `frames` package. Each frame comes in standard, `_large` (2x), and `_xlarge` (3x) variants for high-DPI screens.


### AmigaLike

![](images/gallery_amiga_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.AmigaLike.Image, frames.AmigaLike.Middle)
```

### AmigaLikeLarge

![](images/gallery_amiga_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.AmigaLikeLarge.Image, frames.AmigaLikeLarge.Middle)
```

### AmigaLikeXlarge

![](images/gallery_amiga_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.AmigaLikeXlarge.Image, frames.AmigaLikeXlarge.Middle)
```

### BeosLike

![](images/gallery_beos_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.BeosLike.Image, frames.BeosLike.Middle)
```

### BeosLikeLarge

![](images/gallery_beos_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.BeosLikeLarge.Image, frames.BeosLikeLarge.Middle)
```

### BeosLikeXlarge

![](images/gallery_beos_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.BeosLikeXlarge.Image, frames.BeosLikeXlarge.Middle)
```

### Chains

![](images/gallery_chains.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Chains.Image, frames.Chains.Middle)
```

### ChainsLarge

![](images/gallery_chains_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ChainsLarge.Image, frames.ChainsLarge.Middle)
```

### ChainsXlarge

![](images/gallery_chains_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ChainsXlarge.Image, frames.ChainsXlarge.Middle)
```

### Checkers

![](images/gallery_checkers.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Checkers.Image, frames.Checkers.Middle)
```

### CheckersLarge

![](images/gallery_checkers_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.CheckersLarge.Image, frames.CheckersLarge.Middle)
```

### CheckersXlarge

![](images/gallery_checkers_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.CheckersXlarge.Image, frames.CheckersXlarge.Middle)
```

### ChinaPattern

![](images/gallery_china_pattern.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ChinaPattern.Image, frames.ChinaPattern.Middle)
```

### ChinaPatternLarge

![](images/gallery_china_pattern_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ChinaPatternLarge.Image, frames.ChinaPatternLarge.Middle)
```

### ChinaPatternXlarge

![](images/gallery_china_pattern_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ChinaPatternXlarge.Image, frames.ChinaPatternXlarge.Middle)
```

### Dots

![](images/gallery_dots.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Dots.Image, frames.Dots.Middle)
```

### DotsLarge

![](images/gallery_dots_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.DotsLarge.Image, frames.DotsLarge.Middle)
```

### DotsXlarge

![](images/gallery_dots_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.DotsXlarge.Image, frames.DotsXlarge.Middle)
```

### FantasyStone

![](images/gallery_fantasy_stone.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.FantasyStone.Image, frames.FantasyStone.Middle)
```

### FantasyStoneLarge

![](images/gallery_fantasy_stone_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.FantasyStoneLarge.Image, frames.FantasyStoneLarge.Middle)
```

### FantasyStoneXlarge

![](images/gallery_fantasy_stone_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.FantasyStoneXlarge.Image, frames.FantasyStoneXlarge.Middle)
```

### Floral

![](images/gallery_floral.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Floral.Image, frames.Floral.Middle)
```

### FloralLarge

![](images/gallery_floral_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.FloralLarge.Image, frames.FloralLarge.Middle)
```

### FloralXlarge

![](images/gallery_floral_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.FloralXlarge.Image, frames.FloralXlarge.Middle)
```

### Gold

![](images/gallery_gold.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Gold.Image, frames.Gold.Middle)
```

### GoldLarge

![](images/gallery_gold_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.GoldLarge.Image, frames.GoldLarge.Middle)
```

### GoldXlarge

![](images/gallery_gold_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.GoldXlarge.Image, frames.GoldXlarge.Middle)
```

### Hearts

![](images/gallery_hearts.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Hearts.Image, frames.Hearts.Middle)
```

### HeartsLarge

![](images/gallery_hearts_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.HeartsLarge.Image, frames.HeartsLarge.Middle)
```

### HeartsXlarge

![](images/gallery_hearts_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.HeartsXlarge.Image, frames.HeartsXlarge.Middle)
```

### MacClassicLike

![](images/gallery_mac_classic_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MacClassicLike.Image, frames.MacClassicLike.Middle)
```

### MacClassicLikeLarge

![](images/gallery_mac_classic_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MacClassicLikeLarge.Image, frames.MacClassicLikeLarge.Middle)
```

### MacClassicLikeXlarge

![](images/gallery_mac_classic_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MacClassicLikeXlarge.Image, frames.MacClassicLikeXlarge.Middle)
```

### MacosxLike

![](images/gallery_macosx_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MacosxLike.Image, frames.MacosxLike.Middle)
```

### MacosxLikeLarge

![](images/gallery_macosx_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MacosxLikeLarge.Image, frames.MacosxLikeLarge.Middle)
```

### MacosxLikeXlarge

![](images/gallery_macosx_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MacosxLikeXlarge.Image, frames.MacosxLikeXlarge.Middle)
```

### Metal

![](images/gallery_metal.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Metal.Image, frames.Metal.Middle)
```

### MetalLarge

![](images/gallery_metal_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MetalLarge.Image, frames.MetalLarge.Middle)
```

### MetalXlarge

![](images/gallery_metal_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MetalXlarge.Image, frames.MetalXlarge.Middle)
```

### MwmLike

![](images/gallery_mwm_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MwmLike.Image, frames.MwmLike.Middle)
```

### MwmLikeLarge

![](images/gallery_mwm_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MwmLikeLarge.Image, frames.MwmLikeLarge.Middle)
```

### MwmLikeXlarge

![](images/gallery_mwm_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.MwmLikeXlarge.Image, frames.MwmLikeXlarge.Middle)
```

### NextLike

![](images/gallery_next_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.NextLike.Image, frames.NextLike.Middle)
```

### NextLikeLarge

![](images/gallery_next_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.NextLikeLarge.Image, frames.NextLikeLarge.Middle)
```

### NextLikeXlarge

![](images/gallery_next_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.NextLikeXlarge.Image, frames.NextLikeXlarge.Middle)
```

### Rainbow

![](images/gallery_rainbow.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Rainbow.Image, frames.Rainbow.Middle)
```

### RainbowLarge

![](images/gallery_rainbow_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.RainbowLarge.Image, frames.RainbowLarge.Middle)
```

### RainbowXlarge

![](images/gallery_rainbow_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.RainbowXlarge.Image, frames.RainbowXlarge.Middle)
```

### Ridge

![](images/gallery_ridge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Ridge.Image, frames.Ridge.Middle)
```

### RidgeLarge

![](images/gallery_ridge_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.RidgeLarge.Image, frames.RidgeLarge.Middle)
```

### RidgeXlarge

![](images/gallery_ridge_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.RidgeXlarge.Image, frames.RidgeXlarge.Middle)
```

### ScifiTech

![](images/gallery_scifi_tech.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ScifiTech.Image, frames.ScifiTech.Middle)
```

### ScifiTechLarge

![](images/gallery_scifi_tech_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ScifiTechLarge.Image, frames.ScifiTechLarge.Middle)
```

### ScifiTechXlarge

![](images/gallery_scifi_tech_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.ScifiTechXlarge.Image, frames.ScifiTechXlarge.Middle)
```

### SignConstruction

![](images/gallery_sign_construction.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignConstruction.Image, frames.SignConstruction.Middle)
```

### SignConstructionLarge

![](images/gallery_sign_construction_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignConstructionLarge.Image, frames.SignConstructionLarge.Middle)
```

### SignConstructionXlarge

![](images/gallery_sign_construction_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignConstructionXlarge.Image, frames.SignConstructionXlarge.Middle)
```

### SignStreet

![](images/gallery_sign_street.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignStreet.Image, frames.SignStreet.Middle)
```

### SignStreetLarge

![](images/gallery_sign_street_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignStreetLarge.Image, frames.SignStreetLarge.Middle)
```

### SignStreetXlarge

![](images/gallery_sign_street_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignStreetXlarge.Image, frames.SignStreetXlarge.Middle)
```

### SignWarning

![](images/gallery_sign_warning.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignWarning.Image, frames.SignWarning.Middle)
```

### SignWarningLarge

![](images/gallery_sign_warning_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignWarningLarge.Image, frames.SignWarningLarge.Middle)
```

### SignWarningXlarge

![](images/gallery_sign_warning_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.SignWarningXlarge.Image, frames.SignWarningXlarge.Middle)
```

### Waves

![](images/gallery_waves.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Waves.Image, frames.Waves.Middle)
```

### WavesLarge

![](images/gallery_waves_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WavesLarge.Image, frames.WavesLarge.Middle)
```

### WavesXlarge

![](images/gallery_waves_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WavesXlarge.Image, frames.WavesXlarge.Middle)
```

### Win31Like

![](images/gallery_win31_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Win31Like.Image, frames.Win31Like.Middle)
```

### Win31LikeLarge

![](images/gallery_win31_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Win31LikeLarge.Image, frames.Win31LikeLarge.Middle)
```

### Win31LikeXlarge

![](images/gallery_win31_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Win31LikeXlarge.Image, frames.Win31LikeXlarge.Middle)
```

### Win95Like

![](images/gallery_win95_like.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Win95Like.Image, frames.Win95Like.Middle)
```

### Win95LikeLarge

![](images/gallery_win95_like_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Win95LikeLarge.Image, frames.Win95LikeLarge.Middle)
```

### Win95LikeXlarge

![](images/gallery_win95_like_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Win95LikeXlarge.Image, frames.Win95LikeXlarge.Middle)
```

### WindowFuture

![](images/gallery_window_future.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowFuture.Image, frames.WindowFuture.Middle)
```

### WindowFutureLarge

![](images/gallery_window_future_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowFutureLarge.Image, frames.WindowFutureLarge.Middle)
```

### WindowFutureXlarge

![](images/gallery_window_future_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowFutureXlarge.Image, frames.WindowFutureXlarge.Middle)
```

### WindowGlass

![](images/gallery_window_glass.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowGlass.Image, frames.WindowGlass.Middle)
```

### WindowGlassLarge

![](images/gallery_window_glass_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowGlassLarge.Image, frames.WindowGlassLarge.Middle)
```

### WindowGlassXlarge

![](images/gallery_window_glass_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowGlassXlarge.Image, frames.WindowGlassXlarge.Middle)
```

### WindowPaper

![](images/gallery_window_paper.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowPaper.Image, frames.WindowPaper.Middle)
```

### WindowPaperLarge

![](images/gallery_window_paper_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowPaperLarge.Image, frames.WindowPaperLarge.Middle)
```

### WindowPaperXlarge

![](images/gallery_window_paper_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowPaperXlarge.Image, frames.WindowPaperXlarge.Middle)
```

### WindowRetro

![](images/gallery_window_retro.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowRetro.Image, frames.WindowRetro.Middle)
```

### WindowRetroLarge

![](images/gallery_window_retro_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowRetroLarge.Image, frames.WindowRetroLarge.Middle)
```

### WindowRetroXlarge

![](images/gallery_window_retro_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WindowRetroXlarge.Image, frames.WindowRetroXlarge.Middle)
```

### Wood

![](images/gallery_wood.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.Wood.Image, frames.Wood.Middle)
```

### WoodLarge

![](images/gallery_wood_large.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WoodLarge.Image, frames.WoodLarge.Middle)
```

### WoodXlarge

![](images/gallery_wood_xlarge.png)

```go
import "github.com/arran4/golang-frame/frames"

// Use it directly
fr := frame.NewFrame(destRect, frames.WoodXlarge.Image, frames.WoodXlarge.Middle)
```

