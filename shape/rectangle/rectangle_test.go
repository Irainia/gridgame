package shape_test

import (
	"testing"

	shape "github.com/Irainia/gridgame/shape/rectangle"
)

func TestNewRectangleRowIsNegative(t *testing.T) {
	_, errActual := shape.NewRectangle(-1, 1)
	errExpected := "row should be more than zero"

	if errActual == nil {
		t.Errorf("test_error: %s", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual)
	}
}

func TestNewRectangleRowIsZero(t *testing.T) {
	_, errActual := shape.NewRectangle(0, 1)
	errExpected := "row should be more than zero"

	if errActual == nil {
		t.Errorf("test_error: %s", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual)
	}
}

func TestNewRectangleColumnIsNegative(t *testing.T) {
	_, errActual := shape.NewRectangle(1, -1)
	errExpected := "column should be more than zero"

	if errActual == nil {
		t.Errorf("test_error: %s", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual)
	}
}

func TestNewRectangleColumnIsZero(t *testing.T) {
	_, errActual := shape.NewRectangle(1, 0)
	errExpected := "column should be more than zero"

	if errActual == nil {
		t.Errorf("test_error: %s", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual)
	}
}

func TestOccupiedAreaNotInitialized(t *testing.T) {
	rectangle := shape.Rectangle{}
	occupied := rectangle.GetOccupiedArea()
	if occupied != nil {
		t.Errorf("occupied should be nil")
	}
}

func TestCalculateAreaNotInitialized(t *testing.T) {
	rectangle := shape.Rectangle{}
	area := rectangle.CalculateArea()
	if area != -1 {
		t.Errorf("area should be -1")
	}
}

func TestGetSizeNotInitialized(t *testing.T) {
	rectangle := shape.Rectangle{}
	row, column := rectangle.GetSize()

	if !(row == -1 && column == -1) {
		t.Error("row and column should be -1")
	}
}

func TestNewRectangleCorrectlyInitialized(t *testing.T) {
	_, err := shape.NewRectangle(1, 1)
	if err != nil {
		t.Errorf("test_error: %s", err.Error())
		return
	}
}

func TestGetSizeCorrectlyInitialized(t *testing.T) {
	rect, err := shape.NewRectangle(3, 3)
	if err != nil {
		t.Errorf("test_error: %s", err.Error())
		return
	}

	row, column := rect.GetSize()
	if !(row == 3 && column == 3) {
		t.Error("test_error: row and column should be 3")
	}
}

func TestGetOccupiedAreaCorrectlyInitialized(t *testing.T) {
	rect, err := shape.NewRectangle(3, 3)
	if err != nil {
		t.Errorf("test_error: %s", err.Error())
		return
	}

	expectedOccupied := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}
	actualOccupied := rect.GetOccupiedArea()
	if actualOccupied == nil {
		t.Error("test_error: expected value")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if actualOccupied[i][j] != expectedOccupied[i][j] {
				t.Error("test_error: value doesn't match")
			}
		}
	}
}

func TestCalculateAreaCorrectlyInitialized(t *testing.T) {
	rect, err := shape.NewRectangle(3, 3)
	if err != nil {
		t.Errorf("test_error: %s", err.Error())
		return
	}

	actualArea := rect.CalculateArea()
	if actualArea == -1 {
		t.Error("test_error: area should be more than zero")
	}

	if actualArea != 9 {
		t.Errorf("expected: 9 - actual: %d", actualArea)
	}
}
