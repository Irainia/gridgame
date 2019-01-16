package shape_test

import (
	"testing"

	shape "github.com/irainia/gridgame/shape/triangle"
)

func TestRightButtomTriangleRowIsNegative(t *testing.T) {
	_, error := shape.NewRightButtomTriangle(-1, 1)
	expected := "row should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestRightButtomTriangleRowIsZero(t *testing.T) {
	_, error := shape.NewRightButtomTriangle(0, 1)
	expected := "row should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestRightButtomTriangleColumnIsNegative(t *testing.T) {
	_, error := shape.NewRightButtomTriangle(1, -1)
	expected := "column should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestRightButtomTriangleColumnIsZero(t *testing.T) {
	_, error := shape.NewRightButtomTriangle(1, 0)
	expected := "column should be more than zero"
	if error == nil {
		t.Errorf(expected)
		return
	}

	if error.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, error.Error())
	}
}

func TestRightButtomTriangleExactShape(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewRightButtomTriangle(row, column)

	actual := triangle.GetOccupiedArea()
	expected := [][]bool{
		{false, false, true},
		{false, true, true},
		{true, true, true},
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

	if !isSame {
		t.Errorf("triangle should be the same")
	}
}

func TestRightButtomTriangleWrongShape(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewRightButtomTriangle(row, column)

	actual := triangle.GetOccupiedArea()
	expected := [][]bool{
		{true, false, false},
		{true, true, false},
		{true, true, true},
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

func TestRightButtomTriangleUninitializedOccupied(t *testing.T) {
	triangle := shape.Triangle{}
	actual := triangle.GetOccupiedArea()
	if actual != nil {
		t.Errorf("expected: nil - actual: %v", actual)
	}
}

func TestRightButtomTriangleExactArea(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewRightButtomTriangle(row, column)

	actual := triangle.CalculateArea()
	expected := 6
	if actual != expected {
		t.Errorf("expected: %d - actual: %d", expected, actual)
	}
}

func TestRightButtomTriangleUninitializedArea(t *testing.T) {
	triangle := shape.Triangle{}
	actual := triangle.CalculateArea()
	expected := -1
	if actual != expected {
		t.Errorf("expected: %d - actual: %d", expected, actual)
	}
}

func TestRightButtomTriangleExactSize(t *testing.T) {
	row, column := 3, 3
	triangle, _ := shape.NewRightButtomTriangle(row, column)

	actualRow, actualColumn := triangle.GetSize()
	if !(row == actualRow && column == actualColumn) {
		t.Errorf("expected: (%d,%d) - actual: (%d,%d)", row, column, actualRow, actualRow)
	}
}

func TestRightButtomTriangleUninitializedSize(t *testing.T) {
	triangle := shape.Triangle{}
	actualRow, actualColumn := triangle.GetSize()
	if !(actualRow == -1 && actualColumn == -1) {
		t.Errorf("expected: (-1,-1) - actual: (%d,%d)", actualRow, actualRow)
	}
}
