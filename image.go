package main

import (
	"image"
	"image/color"
	"image/png"
        "os"
        "log"
)

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

func getBlue(distance float32) color.RGBA {
    return color.RGBA{
        R: 0,
        G: 0,
        B: uint8(distance * float32(255)),
        A: 255,
    }
}

