package shape

// Triangle will save rectangle triangle
type Triangle struct {
	occupiedArea [][]bool
	row, column  *int
}

func generateLeftButtomTriangle(row, column int) [][]bool {
	occupied := make([][]bool, row)
	for i := 0; i < row; i++ {
		occupied[i] = make([]bool, column)
	}

	counter := 0.0
	fraction := float64(column) / float64(row)
	var tempNumber int
	for i := 0; i < row; i++ {
		counter += fraction
		tempNumber = int(counter)
		for j := 0; j < tempNumber; j++ {
			occupied[i][j] = true
		}
	}

	minRow, maxRow := -1, -1
	minColumn, maxColumn := 4, -1

	for i := 0; i < row; i++ {
		isrow := false
		tpmin, tpmax := column, -1
		iscolumn := false
		for j := 0; j < column; j++ {
			if !iscolumn {
				if occupied[i][j] {
					tpmin = j
					iscolumn = true
				}
			}

			if occupied[i][j] {
				tpmax = j
			}

			if !isrow {
				isrow = occupied[i][j]
			}
		}

		if tpmin < minColumn {
			minColumn = tpmin
		}
		if tpmax > maxColumn {
			maxColumn = tpmax
		}

		if isrow {
			if minRow < 0 {
				minRow = i
				maxRow = i
			}

			maxRow++
		}
	}
	maxRow--

	r, c := maxRow-minRow+1, maxColumn-minColumn+1
	output := make([][]bool, r)
	for i := 0; i < r; i++ {
		output[i] = make([]bool, c)
		for j := 0; j < c; j++ {
			output[i][j] = occupied[i+minRow][j+minColumn]
		}
	}

	return output
}

func (t *Triangle) reflectOnHorizontalAxis() error {
	occupied := t.occupiedArea
	output := make([][]bool, *t.row)
	for i := 0; i < *t.row; i++ {
		output[i] = make([]bool, *t.column)
	}
	for i := 0; i < *t.row; i++ {
		for j := 0; j < *t.column; j++ {
			output[i][j] = occupied[*t.row-(i+1)][j]
		}
	}

	t.occupiedArea = output
	return nil
}

func (t *Triangle) reflectOnVerticalAxis() error {
	occupied := t.occupiedArea
	output := make([][]bool, *t.row)
	for i := 0; i < *t.row; i++ {
		output[i] = make([]bool, *t.column)
	}
	for i := 0; i < *t.row; i++ {
		for j := 0; j < *t.column; j++ {
			output[i][j] = occupied[i][*t.column-(j+1)]
		}
	}

	t.occupiedArea = output
	return nil
}

// GetOccupiedArea will genereate
func (t Triangle) GetOccupiedArea() [][]bool {
	if t.row == nil || t.column == nil {
		return nil
	}

	output := make([][]bool, *t.row)
	for i := 0; i < *t.row; i++ {
		output[i] = make([]bool, *t.column)
		copy(output[i], t.occupiedArea[i])
	}

	return output
}

// CalculateArea will calculate area
func (t Triangle) CalculateArea() int {
	if t.row == nil || t.column == nil {
		return -1
	}

	occupied := t.occupiedArea
	output := 0
	for i := 0; i < *t.row; i++ {
		for j := 0; j < *t.column; j++ {
			if occupied[i][j] {
				output++
			}
		}
	}

	return output
}

// GetSize will return row, column
func (t Triangle) GetSize() (int, int) {
	if t.row == nil || t.column == nil {
		return -1, -1
	}
	return *t.row, *t.column
}
