package shape

import "errors"

// Rectangle will manage area of the rectangle
type Rectangle struct {
	occupiedArea [][]bool
	row, column  *int
}

// NewRectangle will initialize new Rectangle
func NewRectangle(row, column int) (*Rectangle, error) {
	if row <= 0 {
		return nil, errors.New("row should be more than zero")
	}
	if column <= 0 {
		return nil, errors.New("column should be more than zero")
	}

	rect := Rectangle{}
	rect.row = &row
	rect.column = &column
	occupied := generateRectangle(row, column)
	rect.occupiedArea = occupied

	return &rect, nil
}

func generateRectangle(row, column int) [][]bool {
	output := make([][]bool, row)
	for i := 0; i < row; i++ {
		output[i] = make([]bool, column)
		for j := 0; j < column; j++ {
			output[i][j] = true
		}
	}

	return output
}

// GetOccupiedArea will return occupied area
func (r Rectangle) GetOccupiedArea() [][]bool {
	if r.row == nil || r.column == nil {
		return nil
	}

	output := make([][]bool, *r.row)
	for i := 0; i < *r.row; i++ {
		output[i] = make([]bool, *r.column)
		copy(output[i], r.occupiedArea[i])
	}

	return output
}

// GetSize will return row and column of the rectangle
func (r Rectangle) GetSize() (int, int) {
	if r.row == nil || r.column == nil {
		return -1, -1
	}

	return *r.row, *r.column
}

// CalculateArea will return number of area of the rectangle
func (r Rectangle) CalculateArea() int {
	if r.row == nil || r.column == nil {
		return -1
	}

	counter := 0
	for i := 0; i < *r.row; i++ {
		for j := 0; j < *r.column; j++ {
			if r.occupiedArea[i][j] {
				counter++
			}
		}
	}

	return counter
}
