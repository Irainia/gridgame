package circle

import (
	"errors"
	"math"
)

// Circle will be the type of circle
type Circle struct {
	occupiedArea [][]bool
	row, column  *int
}

// NewCircle will create a new circle with `d` is diameter
func NewCircle(d int) (*Circle, error) {
	if d <= 0 {
		return nil, errors.New("d should be more than `0`")
	}

	occupied := generateNewCircle(d)
	row := len(occupied)
	column := len(occupied[0])

	output := Circle{
		row:          &row,
		column:       &column,
		occupiedArea: occupied,
	}

	return &output, nil
}

func generateNewCircle(d int) [][]bool {
	r := d / 2
	circle := make([][]bool, r*2)
	for i := 0; i < r*2; i++ {
		circle[i] = make([]bool, r*2)
	}

	for i := 0; i < r; i++ {
		j := int(math.Pow(math.Pow(float64(r), 2)-math.Pow(float64(i), 2), 0.5)) - 1
		for k := r - j; k < j+r; k++ {
			circle[i+r][k] = true
			circle[r-i][k] = true
		}
	}

	rowMin, rowMax := 2*r, -1
	colMin, colMax := 2*r, -1
	for i := 0; i < 2*r; i++ {
		for j := 0; j < 2*r; j++ {
			if circle[i][j] {
				if i < rowMin {
					rowMin = i
				}
				if i > rowMax {
					rowMax = i
				}
				if j < colMin {
					colMin = j
				}
				if j > colMax {
					colMax = j
				}
			}
		}
	}

	row, column := rowMax-rowMin+1, colMax-colMin+1
	output := make([][]bool, row)
	for i := 0; i < row; i++ {
		output[i] = make([]bool, column)
	}

	for i := rowMin; i < rowMax+1; i++ {
		for j := colMin; j < colMax+1; j++ {
			output[i-rowMin][j-colMin] = circle[i][j]
		}
	}
	return output
}

// GetOccupiedArea will get the occupied area
func (c Circle) GetOccupiedArea() [][]bool {
	if c.row == nil || c.column == nil {
		return nil
	}

	output := make([][]bool, *c.row)
	for i := 0; i < *c.row; i++ {
		output[i] = make([]bool, *c.column)
		copy(output[i], c.occupiedArea[i])
	}

	return output
}

// GetSize will get the size of the circle
func (c Circle) GetSize() (int, int) {
	if c.row == nil || c.column == nil {
		return -1, -1
	}

	return *c.row, *c.column
}

// CalculateArea the are occupied by circle
func (c Circle) CalculateArea() int {
	if c.row == nil || c.column == nil {
		return -1
	}

	occupied := c.occupiedArea
	output := 0
	for i := 0; i < *c.row; i++ {
		for j := 0; j < *c.column; j++ {
			if occupied[i][j] {
				output++
			}
		}
	}

	return output
}
