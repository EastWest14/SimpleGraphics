package renderer

import (
	"SimpleGraphics/bitmap"
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

//Render writes a PNG image file witht the specified name.
func (r *Renderer) Render(filename string, bitmap *bitmap.Bitmap) error {
	image, err := createImage(bitmap)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("image.png", image, 0644)
	if err != nil {
		return err
	}

	return nil
}

func createImage(bitmap *bitmap.Bitmap) ([]byte, error) {
	m := image.NewRGBA(image.Rect(0, 0, bitmap.Width(), bitmap.Height()))
	for y := 0; y < bitmap.Height(); y++ {
		for x := 0; x < bitmap.Width(); x++ {
			pixel := bitmap.PixelAt(x, y)
			m.Set(x, y, color.RGBA{pixel.Red(), pixel.Green(), pixel.Blue(), 255})
		}
	}

	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
