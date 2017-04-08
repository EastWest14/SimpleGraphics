package renderer

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"SimpleGraphics/bitmap"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

//Render writes a PNG image file witht the specified name.
func (r *Renderer) Render(filename string, bitmap *bitmap.Bitmap) error {
	image, err := createImage()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("image.png", image, 0644)
	if err != nil {
		return err
	}

	return nil
}

func createImage() ([]byte, error) {
	m := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			m.Set(x, y, color.RGBA{uint8(x), uint8(y), uint8(100), 255})
		}
	}

	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
