package logics

type Dimension struct {
	height int
	width int
}

func NewDimension(height, width int) Dimension {
	return Dimension{height, width}
}
