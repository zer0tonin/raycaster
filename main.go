package main

import (
	"image/color"
	"log"
	"os"
        "sync"

	"github.com/hschendel/stl"
)

func readSTL() *stl.Solid {
    fileName := os.Args[1]
    solid, err := stl.ReadFile(fileName)
    if err != nil {
        log.Panicf("Error while reading stl file: %v", err)
    }

    solid.ScaleLinearDowntoSizeBox(stl.Vec3{0.80, 0.80, 0.80})
    solid.Rotate(stl.Vec3{0,0,0}, stl.Vec3{0,1,0}, stl.HalfPi + stl.QuarterPi)
    solid.Rotate(stl.Vec3{0,0,0}, stl.Vec3{1,0,0}, stl.QuarterPi)
    solid.Rotate(stl.Vec3{0,0,0}, stl.Vec3{0,0,1}, stl.HalfPi)
    solid.MoveToPositive()
    solid.Translate(stl.Vec3{0.10, 0.10, 0.10})

    return solid
}

func main() {
    width := 200
    height := 200

    solid := readSTL()
    image := newImage(width, height)
    var wg sync.WaitGroup

    for i := range width {
        for j := range height {
            wg.Add(1)

            go func(i, j int) {
                defer wg.Done()
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
            }(i, j)
        }
    }

    wg.Wait()
    writePNG(image)
}
