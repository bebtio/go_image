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

func ScanNext(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}

func LoadppmImage(imagePath string) Image {

	image := Image{}
	imFile, err := os.Open(imagePath)

	if err != nil {
		log.Fatal(err)
	}

	imScanner := bufio.NewScanner(imFile)

	imScanner.Split(bufio.ScanWords)

	// Magic number. Tells you the file format.
	format := ScanNext(imScanner)

	if format != "P3" {
		log.Fatal("File:", imagePath, "is not format is not P3: Portable Bitmap ASCII. Exiting...")
	}

	// The image dimensions.
	numCols, _ := strconv.ParseUint(ScanNext(imScanner), 10, 32)
	numRows, _ := strconv.ParseUint(ScanNext(imScanner), 10, 32)

	image.numCols = uint32(numCols)
	image.numRows = uint32(numRows)

	// The max size field. I'm not using it here.
	pixels := make([]Pixel, numCols*numRows)

	// Loop over the pixels and get each channel (r,g,b).
	numPixels := uint32(numCols) * uint32(numRows)
	colorChan := uint64(0)
	for index := uint32(0); index < numPixels; index++ {

		// Get red channel
		colorChan, _ = strconv.ParseUint(ScanNext(imScanner), 10, 8)
		pixels[index].r = uint8(colorChan)

		// Get green channel
		colorChan, _ = strconv.ParseUint(ScanNext(imScanner), 10, 8)
		pixels[index].g = uint8(colorChan)

		// Get blue channel
		colorChan, _ = strconv.ParseUint(ScanNext(imScanner), 10, 8)
		pixels[index].b = uint8(colorChan)
	}

	image.pixels = pixels
	return image
}

func DumpImageAscii(im Image) {
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

	DumpImageAscii(image)
}
