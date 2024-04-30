package main

import "github.com/hschendel/stl"

type ray struct {
    origin stl.Vec3
    direction stl.Vec3
}

func castRay(i, j int) ray {
    return ray{
        origin: stl.Vec3{float32(i), float32(j)},
        direction: stl.Vec3{0, 0, 1},
    }
}

func (r ray) intersect(triangle stl.Triangle) (float32, bool) {
    edge1 := sub(triangle.Vertices[1], triangle.Vertices[0])
    edge2 := sub(triangle.Vertices[2], triangle.Vertices[0])

    crossDirEdge2 := cross(r.direction, edge2)

    determinant := dot(edge1, crossDirEdge2)
    if floatEqual(determinant, 0) {
        return 0, false
    }
    invDet := 1 / determinant

    origMinusVert0 := sub(r.origin, triangle.Vertices[0])

    baryU := dot(origMinusVert0, crossDirEdge2) * invDet
    if baryU < 0 || baryU > 1 {
        return 0, false
    }

    crossOriMinusVert0Edge1 := cross(origMinusVert0, edge1)

    baryV := dot(r.direction, crossOriMinusVert0Edge1) * invDet
    if baryV < 0 || baryU + baryV > 1 {
        return 0, false
    }

    rayT := dot(edge2, crossOriMinusVert0Edge1) * invDet

    return rayT, true
}
