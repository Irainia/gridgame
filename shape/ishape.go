package shape

// IShape will be the template for all shapes
type IShape interface {
	GetOccupiedArea() [][]bool
	GetSize() (int, int)
	CalculateArea() int
}
