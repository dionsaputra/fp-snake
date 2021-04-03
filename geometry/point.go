package geometry

import (
	"github.com/dionsaputra/fp-snake/math"
)

type Point struct {
	x int
	y int
}

func PointOf(x, y int) Point {
	return Point{x, y}
}

func (p Point) Translate(direction Direction) Point {
	return PointOf(p.x+direction.dx, p.y+direction.dy)
}

func (p Point) GetDirection(follower Point) Direction {
	return DirectionOf(math.Sign(p.x-follower.x), math.Sign(p.y-follower.y))
}
