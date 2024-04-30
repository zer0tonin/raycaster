package main

import (
	"image/color"
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

    //solid.ScaleLinearDowntoSizeBox(stl.Vec3{1, 1, 1})

    return solid
}

func main() {
    width := 200
    height := 200

    solid := readSTL()
    image := newImage(width, height)

    for i := range width {
        for j := range height {
            image.Set(i, height-j, color.White)

            distance := float32(1)
            intersects := false
            ray := castRay(
                float32(i) / (float32(width) - 1),
                float32(j) / (float32(height) - 1),
            )
            for _, triangle := range solid.Triangles {
                t, ok := ray.intersect(triangle)
                if ok {
                    if t < distance  {
                        distance = t
                    }
                    intersects = true
                }
            }

            if intersects {
                image.Set(i, height-j, getBlue(distance))
            }
        }
    }

    writePNG(image)
}
