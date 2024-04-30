package main

import (
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hschendel/stl"
)

func readSTL() *stl.Solid {
    fileName := os.Args[1]
    solid, err := stl.ReadFile(fileName)
    if err != nil {
        log.Panicf("Error while reading stl file: %v", err)
    }

    //solid.ScaleLinearDowntoSizeBox(stl.Vec3{1, 1, 1})

    return solid
}

func main() {
    width := 800
    height := 600

    solid := readSTL()
    image := newImage(width, height)

    for i := range width {
        for j := range height {
            image.Set(i, j, color.White)

            minT := float32(math.MaxFloat32)
            ray := castRay(i / width, j / height)
            for _, triangle := range solid.Triangles {
                t, ok := ray.intersect(triangle)
                if ok {
                    if t < minT  {
                        minT = t
                    }
                }
            }

            if minT < float32(math.MaxFloat32) {
                image.Set(i, j, getBlue(minT))
            }
        }
    }

    writePNG(image)
}
