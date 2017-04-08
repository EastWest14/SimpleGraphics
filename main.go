package main

import (
	"fmt"

	"SimpleGraphics/renderer"
)

func main() {
	aRenderer := renderer.NewRenderer()
	err := aRenderer.Render("image.png", nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Image written!")
}
