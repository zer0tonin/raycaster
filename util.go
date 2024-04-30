package main

import (
    "math"
)

func floatEqual(x, y float32) bool {
    return math.Nextafter32(x, y) == x
}
