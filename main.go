package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"bytes"
	"io/ioutil"
)

func main() {
	image, err := createImage()
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile("image.png", image, 0644)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Image written!")

}

func createImage() ([]byte, error) {
	m := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for y := 0; y < 256; y++ {
 		for x := 0; x < 256; x++ {
			m.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}

	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}