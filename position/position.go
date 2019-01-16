package position

import "errors"

// Position will hold position in term of row, column
type Position struct {
	row, column *int
}

// NewPosition will initialize Position
func NewPosition(row, column int) (*Position, error) {
	if row < 0 {
		return nil, errors.New("row should be at least zero")
	}
	if column < 0 {
		return nil, errors.New("column should be at least zero")
	}

	output := Position{}
	output.row = &row
	output.column = &column

	return &output, nil
}

// GetPosition will return (row, column)
func (p Position) GetPosition() ([]int, error) {
	if p.row == nil || p.column == nil {
		return nil, errors.New("cannot retrieve row or column value")
	}

	return []int{*p.row, *p.column}, nil
}
