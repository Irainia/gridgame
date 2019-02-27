package circle_test

import (
	"testing"

	"github.com/Irainia/gridgame/shape/circle"
)

func TestNewCircleSuccess(t *testing.T) {
	output, err := circle.NewCircle(6)
	if err != nil {
		t.Errorf("expected: nil - actual: %v", err)
		return
	}

	expected := [][]bool{
		{false, true, true, false},
		{false, true, true, false},
		{true, true, true, true},
		{false, true, true, false},
		{false, true, true, false},
	}

	actual := output.GetOccupiedArea()
	row, column := output.GetSize()
	isSuccess := true
	for i := 0; i < row; i++ {
		if isSuccess {
			for j := 0; j < column; j++ {
				if actual[i][j] != expected[i][j] {
					isSuccess = false
					break
				}
			}
		}
	}

	if !isSuccess {
		t.Error("expected: success - actual: fail")
	}
}

func TestNewCircleDiameterLessThanZero(t *testing.T) {
	_, err := circle.NewCircle(-1)
	if err == nil {
		t.Error("expected: not nil - actual: nil")
	}
}

func TestGetOccupiedAreaFail(t *testing.T) {
	cr := circle.Circle{}
	actual := cr.GetOccupiedArea()
	if actual != nil {
		t.Errorf("expected: nil - actual: %v", actual)
	}
}

func TestGetSizeFail(t *testing.T) {
	cr := circle.Circle{}
	row, column := cr.GetSize()
	if !(row == -1 && column == -1) {
		t.Errorf("expected: {row: -1, column: -1} - actual: {row: %d, column: %d}", row, column)
	}
}

func TestCalculateAreaFail(t *testing.T) {
	cr := circle.Circle{}
	area := cr.CalculateArea()
	if area != -1 {
		t.Errorf("expected: -1 - actual: %d", area)
	}
}

func TestCalculateAreaSucces(t *testing.T) {
	output, err := circle.NewCircle(6)
	if err != nil {
		t.Errorf("expected: nil - actual: %v", err)
		return
	}

	expected := 12
	actual := output.CalculateArea()

	if actual != expected {
		t.Errorf("expected: %d - actual: %d", expected, actual)
	}
}
