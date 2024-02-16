package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pixel struct {
	r, g, b uint8
}

type Image struct {
	numRows uint32
	numCols uint32
	pixels  []Pixel
}

func (im Image) GetPixel(row uint32, col uint32) Pixel {

	index := im.numCols*row + col

	return im.pixels[index]
}

func LoadppmImage(imagePath string) Image {

	image := Image{}
	imFile, err := os.Open(imagePath)

	if err != nil {
		log.Fatal(err)
	}

	imScanner := bufio.NewScanner(imFile)

	imScanner.Split(bufio.ScanWords)

	imScanner.Scan()
	imScanner.Scan()

	numCols, _ := strconv.ParseUint(imScanner.Text(), 10, 32)
	numRows, _ := strconv.ParseUint(imScanner.Text(), 10, 32)

	image.numCols = uint32(numCols)
	image.numRows = uint32(numRows)

	imScanner.Scan()

	pixels := make([]Pixel, numCols*numRows)

	numPixels := uint32(numCols) * uint32(numRows)
	colorChan := uint64(0)
	for index := uint32(0); index < numPixels; index++ {

		// Get red channel
		imScanner.Scan()
		colorChan, _ = strconv.ParseUint(imScanner.Text(), 10, 8)
		pixels[index].r = uint8(colorChan)

		// Get green channel
		imScanner.Scan()
		colorChan, _ = strconv.ParseUint(imScanner.Text(), 10, 8)
		pixels[index].g = uint8(colorChan)

		// Get blue channel
		imScanner.Scan()
		colorChan, _ = strconv.ParseUint(imScanner.Text(), 10, 8)
		pixels[index].b = uint8(colorChan)
	}

	image.pixels = pixels
	return image
}

func dumpImageAscii(im Image) {
	for r := uint32(0); r < im.numRows; r++ {
		for c := uint32(0); c < im.numCols; c++ {
			p := im.GetPixel(r, c)

			fmt.Printf("%d %d %d   ", p.r, p.g, p.b)
		}
		fmt.Println("")
	}
}

func main() {

	image := LoadppmImage("image.ppm")

	dumpImageAscii(image)
}
