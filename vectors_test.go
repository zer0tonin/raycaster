package main

import (
	"testing"

	"github.com/hschendel/stl"
)

func TestSub(t *testing.T) {
    tests := []struct{
        name string
        v1 stl.Vec3
        v2 stl.Vec3
        expected stl.Vec3
    }{
        {
            name: "0 vectors",
            v1: stl.Vec3{0, 0, 0},
            v2: stl.Vec3{0, 0, 0},
            expected: stl.Vec3{0, 0, 0},
        },
        {
            name: "positive values",
            v1: stl.Vec3{3, 5, 0},
            v2: stl.Vec3{1, 2, 3},
            expected: stl.Vec3{2, 3, -3},
        },
        {
            name: "positive and negative values",
            v1: stl.Vec3{3, -5, 0},
            v2: stl.Vec3{1, 2, -3},
            expected: stl.Vec3{2, -7, 3},
        },
    }

    for _, test := range tests {
        actual := sub(test.v1, test.v2)

        for i, a := range actual {
            if !floatEqual(test.expected[i], a) {
                t.Errorf("Incorrect substraction in %s, expected %f got %f", test.name, test.expected[i], a)
            }
        }
    }
} 

func TestDot(t *testing.T) {
    tests := []struct{
        name string
        v1 stl.Vec3
        v2 stl.Vec3
        expected float32
    }{
        {
            name: "0 vectors",
            v1: stl.Vec3{0, 0, 0},
            v2: stl.Vec3{0, 0, 0},
            expected: 0,
        },
        {
            name: "parallel",
            v1: stl.Vec3{1, 1, 1},
            v2: stl.Vec3{2, 2, 2},
            expected: 6,
        },
        {
            name: "perpendicular",
            v1: stl.Vec3{0, 1, 1},
            v2: stl.Vec3{1, 0, 0},
            expected: 0,
        },
    }

    for _, test := range tests {
        actual := dot(test.v1, test.v2)

        if !floatEqual(test.expected, actual) {
            t.Errorf("Incorrect dot product in %s, expected %f got %f", test.name, test.expected, actual)
        }
    }
}

func TestCross(t *testing.T) {
    tests := []struct{
        name string
        v1 stl.Vec3
        v2 stl.Vec3
        expected stl.Vec3
    }{
        {
            name: "0 vectors",
            v1: stl.Vec3{0, 0, 0},
            v2: stl.Vec3{0, 0, 0},
            expected: stl.Vec3{0, 0, 0},
        },
        {
            name: "parallel",
            v1: stl.Vec3{1, 1, 1},
            v2: stl.Vec3{2, 2, 2},
            expected: stl.Vec3{0, 0, 0},
        },
        {
            name: "perpendicular",
            v1: stl.Vec3{0, 1, 1},
            v2: stl.Vec3{1, 0, 0},
            expected: stl.Vec3{0, 1, -1},
        },
    }

    for _, test := range tests {
        actual := cross(test.v1, test.v2)

        for i, a := range actual {
            if !floatEqual(test.expected[i], a) {
                t.Errorf("Incorrect cross product in %s, expected %f got %f", test.name, test.expected[i], a)
            }
        }
    }
}
