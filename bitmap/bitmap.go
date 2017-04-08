package bitmap

const (
	IMAGE_SIZE = 256
)

type Bitmap struct {
	pixels [][]Pixel
}

func NewBitmap() *Bitmap {
	bm := &Bitmap{}
	bm.pixels = make([][]Pixel, IMAGE_SIZE, IMAGE_SIZE)

	for i, _ := range bm.pixels {
		bm.pixels[i] = make([]Pixel, 256, 256)
	}

	return bm
}

func (b *Bitmap) PixelAt(x, y int) Pixel {
	return b.pixels[x][y]
}

func (b *Bitmap) SetPixelAt(x, y int, p Pixel) {
	b.pixels[x][y] = p
}
