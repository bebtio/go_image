package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"go_image/go_image"
)

func ScanNext(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}

func LoadppmImage(imagePath string) go_image.Image {

	image := go_image.Image{}
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

	image.NumCols = uint32(numCols)
	image.NumRows = uint32(numRows)

	// The max size field. I'm not using it here.
	// May use this later to normalize between 0-255 if max size is greater than 255.
	pixels := make([]go_image.Pixel, numCols*numRows)

	// Loop over the pixels and get each channel (r,g,b).
	numPixels := uint32(numCols) * uint32(numRows)
	colorChan := uint64(0)
	for index := uint32(0); index < numPixels; index++ {

		// Get red channel
		colorChan, _ = strconv.ParseUint(ScanNext(imScanner), 10, 8)
		pixels[index].R = uint8(colorChan)

		// Get green channel
		colorChan, _ = strconv.ParseUint(ScanNext(imScanner), 10, 8)
		pixels[index].G = uint8(colorChan)

		// Get blue channel
		colorChan, _ = strconv.ParseUint(ScanNext(imScanner), 10, 8)
		pixels[index].B = uint8(colorChan)
	}

	image.Pixels = pixels
	return image
}

func DumpImageAscii(im go_image.Image) {
	for r := uint32(0); r < im.NumRows; r++ {
		for c := uint32(0); c < im.NumCols; c++ {
			p := im.GetPixel(r, c)

			fmt.Printf("%d %d %d   ", p.R, p.G, p.B)
		}
		fmt.Println("")
	}
}

func main() {

	image := LoadppmImage("image.ppm")

	DumpImageAscii(image)
}
