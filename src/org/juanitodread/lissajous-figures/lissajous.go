package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black,
	color.RGBA{0, 255, 0, 1}, // Green
	color.RGBA{255, 0, 0, 1}, // Red
	color.RGBA{0, 0, 255, 1}, // Blue
	color.White}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles     = 5     // Number of complete X oscillator revolutions
		resolution = 0.001 // Angular resolution
		size       = 100   // Image canvas covers [-size..+size]
		nFrames    = 64    // Number of animation nFrames
		delay      = 8     // Delay between frames in 10 ms units
		freqY      = 3.0   // Frequency of Y oscillator
	)

	freq := rand.Float64() * freqY
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	colorIndex := 1

	for i := 0; i < nFrames; i++ {
		width := 2*size + 1
		rect := image.Rect(0, 0, width, width)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
			colorIndex++

			if colorIndex == 4 {
				colorIndex = 1
			}

		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
