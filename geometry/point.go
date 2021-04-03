package geometry

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