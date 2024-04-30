# zer0tonin's raycaster

A raycaster writter in Go, that reads STL files and outputs PNG images.
It uses the [Möller–Trumbore intersection algorithm](https://en.wikipedia.org/wiki/M%C3%B6ller%E2%80%93Trumbore_intersection_algorithm) to compute the intersection between rays and triangles.

Usage:
```
raycaster input.stl output.png
```
