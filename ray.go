package main

import "github.com/hschendel/stl"

type ray struct {
    origin *stl.Vec3
    direction *stl.Vec3
}

func castRay(i, j int) ray {
    return ray{}
}

func (r ray) intersect(triangle stl.Triangle) bool {
    return false
}
