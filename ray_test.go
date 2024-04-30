package main

import (
	"testing"

	"github.com/hschendel/stl"
)

func TestIntersect(t *testing.T) {
    tests := []struct{
        name string
        r ray
        triangle stl.Triangle
        intersects bool
        distance float32
    }{
        {
            name: "parallel",
            r: ray{
                origin: stl.Vec3{0.5, 0.5, 0},
                direction: stl.Vec3{0, 0, 1},
            },
            triangle: stl.Triangle{
                Vertices: [3]stl.Vec3{
                    {0, 0, 0},
                    {0, 0, 1},
                    {1, 0, 1},
                },
            },
            intersects: false,
            distance: 0,
        },
        {
            name: "intersects",
            r: ray{
                origin: stl.Vec3{0.75, 0.75, 0},
                direction: stl.Vec3{0, 0, 1},
            },
            triangle: stl.Triangle{
                Vertices: [3]stl.Vec3{
                    {0, 0, 1},
                    {0, 1, 1},
                    {1, 1, 1},
                },
            },
            intersects: true,
            distance: 1,
        },
        {
            name: "doesn't intersect",
            r: ray{
                origin: stl.Vec3{0.75, 0.25, 0},
                direction: stl.Vec3{0, 0, 1},
            },
            triangle: stl.Triangle{
                Vertices: [3]stl.Vec3{
                    {0, 0, 1},
                    {0, 1, 1},
                    {1, 1, 1},
                },
            },
            intersects: false,
            distance: 0,
        },
        {
            name: "intersects (distance 0.5)",
            r: ray{
                origin: stl.Vec3{0.75, 0.75, 0},
                direction: stl.Vec3{0, 0, 1},
            },
            triangle: stl.Triangle{
                Vertices: [3]stl.Vec3{
                    {0, 0, 0.5},
                    {0, 1, 0.5},
                    {1, 1, 0.5},
                },
            },
            intersects: true,
            distance: 0.5,
        },
    }

    for _, test := range tests {
        actual, ok := test.r.intersect(test.triangle)
        if ok != test.intersects {
            t.Errorf("Incorrect intersection in %s, expected %t, got %t", test.name, test.intersects, ok)
            return
        }
        if !floatEqual(actual, test.distance) {
            t.Errorf("Invalid distance in %s, expected %f, got %f", test.name, test.distance, actual)
        }
    }
}
