package main

import (
	"github.com/hschendel/stl"
)

func sub(v1, v2 *stl.Vec3) *stl.Vec3 {
    return &stl.Vec3{
	v1[0] - v2[0],
	v1[1] - v2[1],
	v1[2] - v2[2],
    }
}

func dot(v1, v2 *stl.Vec3) float32 {
    return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func cross(v1, v2 *stl.Vec3) *stl.Vec3 {
    return &stl.Vec3{
	v1[1]*v2[2] - v1[2]*v2[1],
	v1[2]*v2[0] - v1[0]*v2[2],
	v1[0]*v2[1] - v1[1]*v2[0],
    }
}
