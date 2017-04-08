package main

import (
	"SimpleGraphics/camera"
	"SimpleGraphics/renderer"
	"SimpleGraphics/scene"
	"fmt"
)

func main() {
	aScene := scene.NewScene()
	aCamera := camera.NewCamera(aScene)
	aRenderer := renderer.NewRenderer()

	bitmap, err := aCamera.Snapshot()
	if err != nil {
		panic(err.Error())
	}

	err = aRenderer.Render("image.png", bitmap)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Image written!")
}
