package geometry

type Direction struct {
	Vertical   int
	Horizontal int
}

func Left() Direction {
	return Direction{0, -1}
}

func Right() Direction {
	return Direction{0, 1}
}

func Up() Direction {
	return Direction{-1, 0}
}

func Down() Direction {
	return Direction{1, 0}
}

func (d Direction) IsOpposite(direction Direction) bool {
	return (d.Vertical*direction.Vertical)+(d.Horizontal*direction.Horizontal) == -1
}
