package points

type Point struct {
	x int
	y int
}

func PointOf(x, y int) Point {
	return Point{x, y}
}

func (p Point) Translate(dx, dy int) Point {
	return PointOf(p.x+dx, p.y+dy)
}
