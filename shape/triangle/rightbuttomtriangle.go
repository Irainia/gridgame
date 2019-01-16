package shape

import (
	"errors"
)

// NewRightButtomTriangle will create new left triangle
func NewRightButtomTriangle(row, column int) (*Triangle, error) {
	if row <= 0 {
		return nil, errors.New("row should be more than zero")
	}
	if column <= 0 {
		return nil, errors.New("column should be more than zero")
	}

	output := Triangle{}
	output.row = &row
	output.column = &column
	output.occupiedArea = generateLeftButtomTriangle(row, column)
	output.reflectOnVerticalAxis()

	return &output, nil
}
