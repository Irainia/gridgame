package shape_test

import (
	"testing"

	shape "github.com/Irainia/gridgame/shape/triangle"
)

func TestLeftTopTriangleRowIsNegative(t *testing.T) {
	_, error := shape.NewLeftTopTriangle(-1, 1)
	expected := "row should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestLeftTopTriangleRowIsZero(t *testing.T) {
	_, error := shape.NewLeftTopTriangle(0, 1)
	expected := "row should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestLeftTopTriangleColumnIsNegative(t *testing.T) {
	_, error := shape.NewLeftTopTriangle(1, -1)
	expected := "column should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestLeftTopTriangleColumnIsZero(t *testing.T) {
	_, error := shape.NewLeftTopTriangle(1, 0)
	expected := "column should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestLeftTopTriangleExactShape(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewLeftTopTriangle(row, column)

	actual := triangle.GetOccupiedArea()
	expected := [][]bool{{true, true, true},
		{true, true, false},
		{true, false, false}}
	isSame := true

	for i := 0; i < row; i++ {
		if isSame {
			for j := 0; j < column; j++ {
				if actual[i][j] != expected[i][j] {
					isSame = false
					break
				}
			}
		}
	}

	if !isSame {
		t.Errorf("triangle should be the same")
	}
}

func TestLeftTropTriangleWrongShape(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewLeftTopTriangle(row, column)

	actual := triangle.GetOccupiedArea()
	expected := [][]bool{
		{true, true, true},
		{false, true, true},
		{false, false, true},
	}
	isSame := true

	for i := 0; i < row; i++ {
		if isSame {
			for j := 0; j < column; j++ {
				if actual[i][j] != expected[i][j] {
					isSame = false
					break
				}
			}
		}
	}

	if isSame {
		t.Errorf("triangle should be different")
	}
}

func TestLeftTopTriangleUninitializedOccupied(t *testing.T) {
	triangle := shape.Triangle{}
	actual := triangle.GetOccupiedArea()
	if actual != nil {
		t.Errorf("expected: nil - actual: %v", actual)
	}
}

func TestLeftTopTriangleExactArea(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewLeftTopTriangle(row, column)

	actual := triangle.CalculateArea()
	expected := 6
	if actual != expected {
		t.Errorf("expected: %d - actual: %d", expected, actual)
	}
}

func TestLeftTopTriangleUninitializedArea(t *testing.T) {
	triangle := shape.Triangle{}
	actual := triangle.CalculateArea()
	expected := -1
	if actual != expected {
		t.Errorf("expected: %d - actual: %d", expected, actual)
	}
}

func TestLeftTopTriangleExactSize(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewLeftTopTriangle(row, column)

	actualRow, actualColumn := triangle.GetSize()
	if !(row == actualRow && column == actualColumn) {
		t.Errorf("expected: (%d,%d) - actual: (%d,%d)", row, column, actualRow, actualRow)
	}
}

func TestLeftTopTriangleUninitializedSize(t *testing.T) {
	triangle := shape.Triangle{}
	actualRow, actualColumn := triangle.GetSize()
	if !(actualRow == -1 && actualColumn == -1) {
		t.Errorf("expected: (-1,-1) - actual: (%d,%d)", actualRow, actualRow)
	}
}
