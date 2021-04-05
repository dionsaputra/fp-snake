package geometry

type Dimension struct {
	Height int
	Width  int
}

func NewDimension(height, width int) Dimension {
	return Dimension{height, width}
}
