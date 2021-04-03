package geometry

type Direction struct {
	dx int
	dy int
}

func Left() Direction {
	return Direction{-1, 0}
}

func Right() Direction {
	return Direction{1, 0}
}

func Up() Direction {
	return Direction{0, -1}
}

func Down() Direction {
	return Direction{0, 1}
}
