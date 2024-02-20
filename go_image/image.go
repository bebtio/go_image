package go_image

type Pixel struct {
	R, G, B uint8
}

type Image struct {
	NumRows uint32
	NumCols uint32
	Pixels  []Pixel
}

func (im Image) GetPixel(row uint32, col uint32) Pixel {

	index := im.NumCols*row + col

	return im.Pixels[index]
}
