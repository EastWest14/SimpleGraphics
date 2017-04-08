package scene

import (
	"SimpleGraphics/bitmap"
	"fmt"
)

const (
	SCENE_SIZE = 256
)

type Scene struct {
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) ColorAt(x, y int) (bitmap.Pixel, error) {
	if x < 0 || x >= SCENE_SIZE {
		return bitmap.NewPixel(uint8(0), uint8(0), uint8(0), uint8(0)), fmt.Errorf("x is %d, out of range", x)
	}
	if y < 0 || y >= SCENE_SIZE {
		return bitmap.NewPixel(uint8(0), uint8(0), uint8(0), uint8(0)), fmt.Errorf("y is %d, out of range", y)
	}
	return bitmap.NewPixel(uint8(x), uint8(100), uint8(y), uint8(0)), nil
}
