package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"time"
)

// predefined vals
const (
    BOUNDARY float64 = 2
	STEP     float64 = .0005


	//after this value we will break calculations because this means point never diverges from the set
	MANDELBROTMAXITERATION = 50

	//the each how point how to be upscaled this much for image
    IMAGESCALEUP = 1
)

// calculated vals
const (

	//for the sake of simplicity same heigh and width, always will be square
	IMAGESIZE = 2 * (BOUNDARY / STEP) * IMAGESCALEUP

	BOUNDARYSQR = BOUNDARY * BOUNDARY
)

var BLUE = color.RGBA{0, 0, 255, 255}
var GREEN = color.RGBA{0, 255, 0, 255}
var RED = color.RGBA{255, 0, 0, 255}

// The point contans the position and the value for that point
type Point struct {
	X, Y float64
}

func (p *Point) AbsoluteSqr() float64 {
	return p.X*p.X + p.Y*p.Y
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0,int(IMAGESIZE), int(IMAGESIZE)))

	log.Println("Creating modelbrot set", IMAGESIZE)

	// samplePoint := Point {X: 0, Y: 1}
	// loopUntilOut(samplePoint, true)
	// return

    // paintBlock(m, 0, 100,RED, int(IMAGESCALEUP))

	//Now we will calculate the matrix row and columns. The top left corner will be point, (-BOUNDARY, BOUNDARY), one to the right would be (-BOUNDARY + rowStep, BOUNDARY)
	//Sample matrix, BOUNDARY as B
	/*
	   -B,B           | -B + columnStep, B | -B + 2 * columnStep, B | ......
	   -B,B - rowStep | -B + columnsteo, B - rowStep | .......

	*/
	imageX, imageY := 0, 0
	for i := BOUNDARY; i >= -BOUNDARY; i -= STEP {
		for j := -BOUNDARY; j <= BOUNDARY; j += STEP {
			point := Point{
				X: j,
				Y: i,
			}

			iterations := loopUntilOut(point, false)
			color := iterationToColor(iterations)

			paintBlock(m, imageX, imageY, color, int(IMAGESCALEUP))

			imageX += int(IMAGESCALEUP)
		}

		imageY += int(IMAGESCALEUP)
		imageX = 0
	}

	log.Println("finished in memory image, flushing to file...")

	file, err := os.Create("image.png")
	maybePanic(err)
	defer file.Close()

	err = png.Encode(file, m)
	maybePanic(err)
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func paintBlock(m *image.RGBA, imageX int, imageY int, color color.RGBA, blockSize int) {
	for i := 0; i < blockSize; i++ {
		for j := 0; j < blockSize; j++ {
			m.Set(imageX + i,imageY + j, color)
		}
	}
}

func iterationToColor(itrations int) color.RGBA {
	if itrations == MANDELBROTMAXITERATION {
		// Inside the Mandelbrot set: black
		return color.RGBA{0, 0, 0, 255}
	}

    // Generate a hue value with a broader color transition
	hue := 270.0 * (float64(itrations) / float64(MANDELBROTMAXITERATION)) // More saturated blues and purples
	saturation := 1.0
	value := 0.8 + 0.2*math.Sin(float64(itrations)/float64(MANDELBROTMAXITERATION)*math.Pi) // Keep colors bright

	// Convert HSV to RGB
	r, g, b := HSVtoRGB(hue, saturation, value)

	return color.RGBA{r, g, b, 255}

	return color.RGBA{r, g, b, 255}
}

func loopUntilOut(point Point, debug bool) int {
	iterations := 0

	basePoint := point

	for point.AbsoluteSqr() < BOUNDARYSQR && iterations < MANDELBROTMAXITERATION {
		point = point.NextMandelBrotPoint(basePoint)
		if debug {
			log.Println(point)
			time.Sleep(time.Second * 1)
		}

		iterations += 1
	}

	return iterations
}

// calculate next points, z = z^2 + c
// the idea is i^2 = -1, complex number
// z^2 = (x + yi)^2 +c  = (x^2 - y^2, 2xy) + c
func (p *Point) NextMandelBrotPoint(basePoint Point) Point {
	nextPoint := Point{}
	nextPoint.X = (p.X*p.X - p.Y*p.Y) + basePoint.X
	nextPoint.Y = (2 * p.X * p.Y) + basePoint.Y

	return nextPoint
}

// HSVtoRGB converts HSV color values to RGB
func HSVtoRGB(h, s, v float64) (r, g, b uint8) {
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := v - c

	var r1, g1, b1 float64

	switch {
	case h >= 0 && h < 60:
		r1, g1, b1 = c, x, 0
	case h >= 60 && h < 120:
		r1, g1, b1 = x, c, 0
	case h >= 120 && h < 180:
		r1, g1, b1 = 0, c, x
	case h >= 180 && h < 240:
		r1, g1, b1 = 0, x, c
	case h >= 240 && h < 300:
		r1, g1, b1 = x, 0, c
	case h >= 300 && h < 360:
		r1, g1, b1 = c, 0, x
	}

	return uint8((r1 + m) * 255), uint8((g1 + m) * 255), uint8((b1 + m) * 255)
}
