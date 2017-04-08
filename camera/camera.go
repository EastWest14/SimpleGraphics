package camera

import (
	"SimpleGraphics/bitmap"
	"SimpleGraphics/scene"
)

const (
	CAMERA_SIZE = 256
)

type Camera struct {
	attachedScene *scene.Scene
}

func NewCamera(scene *scene.Scene) *Camera {
	return &Camera{attachedScene: scene}
}

func (c *Camera) Snapshot() (*bitmap.Bitmap, error) {
	if c.attachedScene == nil {
		panic("Scene not attached to camera.")
	}

	bitmapImage := bitmap.NewBitmap()
	for i := 0; i < CAMERA_SIZE; i++ {
		for j := 0; j < CAMERA_SIZE; j++ {
			pixel, err := c.attachedScene.ColorAt(i, j)
			if err != nil {
				panic(err.Error())
			}
			bitmapImage.SetPixelAt(i, j, pixel)
		}
	}

	return bitmapImage, nil
}
