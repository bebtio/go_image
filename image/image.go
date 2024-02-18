package image

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
