package position_test

import (
	"testing"

	"github.com/irainia/gridgame/position"
)

func TestNewPositionRowLessThanZero(t *testing.T) {
	_, errActual := position.NewPosition(-1, 0)
	errExpected := "row should be at least zero"
	if errActual == nil {
		t.Errorf("expected: %s - actual: nil", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual.Error())
	}
}

func TestNewPositionColumnLessThanZero(t *testing.T) {
	_, errActual := position.NewPosition(0, -1)
	errExpected := "column should be at least zero"
	if errActual == nil {
		t.Errorf("expected: %s - actual: nil", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual.Error())
	}
}

func TestNewPositionCorrectInit(t *testing.T) {
	_, errActual := position.NewPosition(0, 0)
	if errActual != nil {
		t.Errorf("expected: nil - actual: %s", errActual.Error())
	}
}

func TestGetPositionNilInit(t *testing.T) {
	pos := position.Position{}
	_, errActual := pos.GetPosition()
	errExpected := "cannot retrieve row or column value"
	if errActual == nil {
		t.Errorf("expected: %s - actual: nil", errExpected)
		return
	}

	if errActual.Error() != errExpected {
		t.Errorf("expected: %s - actual: %s", errExpected, errActual.Error())
	}
}

func TestGetPositionCorrectInit(t *testing.T) {
	pos, errActual := position.NewPosition(0, 0)
	if errActual != nil {
		t.Errorf("error_test: %s", errActual.Error())
		return
	}
	_, errActual = pos.GetPosition()
	if errActual != nil {
		t.Errorf("expected: nil - actual: %s", errActual.Error())
	}
}
