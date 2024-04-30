package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hschendel/stl"
)

func readSTL() *stl.Solid {
    fileName := os.Args[1]
    solid, err := stl.ReadFile(fileName)
    if err != nil {
        log.Panicf("Error while reading stl file: %v", err)
    }

    return solid
}

func newImage(width, height int) *image.RGBA {
    rect := image.Rect(0, 0, width, height)
    return image.NewRGBA(rect)
}

func writePNG(image image.Image) {
    output, err := os.Create(os.Args[2])
    if err != nil {
        log.Panicf("Error while creating png file: %v", err)
    }
    defer output.Close()

    if err := png.Encode(output, image); err != nil {
        log.Panicf("Failed to write png: %v", err)
    }
}

func main() {
    width := 800
    height := 600
    white := color.RGBA{
        R: 255,
        G: 255,
        B: 255,
        A: 1,
    }
    blue := color.RGBA{
        R: 40,
        G: 67,
        B: 135,
        A: 1,
    }

    solid := readSTL()
    image := newImage(width, height)

    for i := range width {
        for j := range height {
            ray := castRay(i, j)
            for _, triangle := range solid.Triangles {
                if ray.intersect(triangle) {
                    image.Set(i, j, blue)
                } else {
                    image.Set(i, j, white)
                }
            }
        }
    }

    writePNG(image)
}
