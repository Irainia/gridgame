package board_test

import (
	"testing"

	"github.com/Irainia/gridgame/board"
	"github.com/Irainia/gridgame/position"
	rectangle "github.com/Irainia/gridgame/shape/rectangle"
	triangle "github.com/Irainia/gridgame/shape/triangle"
)

func TestNewGameBoardRowIsNegative(t *testing.T) {
	_, err := board.NewGameBoard(-1, 1)
	expected := "row should be more than zero"

	if err == nil {
		t.Errorf(expected)
		return
	}

	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestNewGameBoardRowIsZero(t *testing.T) {
	_, err := board.NewGameBoard(0, 1)
	expected := "row should be more than zero"

	if err == nil {
		t.Errorf(expected)
		return
	}

	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestNewNewGameBoardIsNegative(t *testing.T) {
	_, err := board.NewGameBoard(1, -1)
	expected := "column should be more than zero"

	if err == nil {
		t.Errorf(expected)
		return
	}

	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestNewGameBoardColumnIsZero(t *testing.T) {
	_, err := board.NewGameBoard(1, 0)
	expected := "column should be more than zero"

	if err == nil {
		t.Errorf(expected)
		return
	}

	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestAddShapeWithBoardNotInitialized(t *testing.T) {
	brd := board.GameBoard{}
	sp, _ := rectangle.NewRectangle(1, 1)
	pos, _ := position.NewPosition(0, 0)
	_, err := brd.AddShape(sp, pos)
	expected := "board is not initialized"
	if err == nil {
		t.Errorf(expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestAddShapeWithShapeIsNil(t *testing.T) {
	brd, _ := board.NewGameBoard(10, 10)
	pos, _ := position.NewPosition(0, 0)
	_, err := brd.AddShape(nil, pos)
	expected := "shape should not be a nil value"
	if err == nil {
		t.Errorf(expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestAddShapeWithShapeNotInitialized(t *testing.T) {
	brd, _ := board.NewGameBoard(10, 10)
	pos, _ := position.NewPosition(0, 0)
	_, err := brd.AddShape(rectangle.Rectangle{}, pos)
	expected := "shape is not initialized"
	if err == nil {
		t.Errorf(expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestAddShapeWithPositionIsNil(t *testing.T) {
	brd, _ := board.NewGameBoard(10, 10)
	sp, _ := rectangle.NewRectangle(2, 2)
	_, err := brd.AddShape(sp, nil)
	expected := "position should not be a nil value"
	if err == nil {
		t.Errorf(expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestAddShapeWithPositionNotInitialized(t *testing.T) {
	brd, _ := board.NewGameBoard(10, 10)
	pos := position.Position{}
	sp, _ := rectangle.NewRectangle(2, 2)
	_, err := brd.AddShape(sp, &pos)
	expected := "position is not initialized"
	if err == nil {
		t.Errorf(expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("expected: %s - actual: %s", expected, err.Error())
	}
}

func TestAddShapePositionExceedBoard(t *testing.T) {
	sp, _ := rectangle.NewRectangle(3, 3)

	brd, _ := board.NewGameBoard(10, 10)
	ps, _ := position.NewPosition(10, 10)
	status, err := brd.AddShape(sp, ps)

	expectedStatus := false
	expectedError := "position exceed board"

	if err == nil {
		t.Errorf("expected: (false, \"%s\") - actual: (%t,nil)",
			expectedError, status)
		return
	}

	if status == expectedStatus && err.Error() != expectedError {
		t.Errorf("expected: (false, \"%s\") - actual: (%t,\"%s\")",
			expectedError, status, err.Error())
	}
}

func TestAddShapeShouldSuccess(t *testing.T) {
	sp, _ := rectangle.NewRectangle(3, 3)

	brd, _ := board.NewGameBoard(10, 10)
	pos, _ := position.NewPosition(0, 0)
	status, err := brd.AddShape(sp, pos)
	expected := "expected: (true, nil)"

	if !status || err != nil {
		t.Errorf(expected)
	}
}

func TestAddShapePositionWithShapeExceedBoard(t *testing.T) {
	sp, _ := rectangle.NewRectangle(3, 3)

	brd, _ := board.NewGameBoard(10, 10)
	ps, _ := position.NewPosition(8, 8)
	status, err := brd.AddShape(sp, ps)

	expectedStatus := true
	expectedError := "position exceed board"

	if err == nil {
		t.Errorf("expected: (true, \"%s\") - actual: (%t,nil)",
			expectedError, status)
		return
	}

	if status != expectedStatus || err.Error() != expectedError {
		t.Errorf("expected: (true, \"%s\") - actual: (%t,\"%s\")",
			expectedError, status, err.Error())
	}
}

func TestAddShapeNotOverlap(t *testing.T) {
	rect, _ := rectangle.NewRectangle(4, 1)
	ltriangle, _ := triangle.NewLeftButtomTriangle(3, 3)
	brd, _ := board.NewGameBoard(4, 4)
	pos := position.Position{}
	brd.AddShape(rect, &pos)

	ps, _ := position.NewPosition(0, 1)
	status, _ := brd.AddShape(ltriangle, ps)
	if !status {
		t.Errorf("adding shape should success")
	}
}

func TestAddShapeOverlap(t *testing.T) {
	ltriangle, _ := triangle.NewLeftButtomTriangle(3, 3)
	rect, _ := rectangle.NewRectangle(4, 1)

	brd, _ := board.NewGameBoard(4, 4)
	pos, _ := position.NewPosition(0, 0)
	brd.AddShape(ltriangle, pos)
	pos, _ = position.NewPosition(0, 0)
	status, err := brd.AddShape(rect, pos)

	expectedStatus := false
	expectedError := "overlap adding shape"
	if err == nil {
		t.Errorf("expected: (false, \"%s\") - actual: (%t,nil)",
			expectedError, status)
		return
	}

	if status != expectedStatus && err.Error() != expectedError {
		t.Errorf("expected: (false, \"%s\") - actual: (%t,\"%s\")",
			expectedError, status, err.Error())
	}
}

func TestGetOccupiedAreaFail(t *testing.T) {
	brd := board.GameBoard{}
	occ := brd.GetOccupiedArea()
	if occ != nil {
		t.Errorf("expected: nil - actuL: %v", occ)
	}
}
func TestGetOccupiedAreaEmpty(t *testing.T) {
	brd, _ := board.NewGameBoard(3, 3)
	actual := brd.GetOccupiedArea()
	expected := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}

	row, column := brd.GetSize()
	isSuccess := true
	for i := 0; i < row; i++ {
		if isSuccess {
			for j := 0; j < column; j++ {
				isSuccess = actual[i][j] == expected[i][j]
				if !isSuccess {
					break
				}
			}
		}
	}

	if !isSuccess {
		t.Errorf("board should be empty")
	}
}

func TestGetOccupiedAreaNotEmpty(t *testing.T) {
	sp, _ := triangle.NewLeftButtomTriangle(3, 3)
	brd, _ := board.NewGameBoard(4, 4)
	pos, _ := position.NewPosition(0, 0)
	brd.AddShape(sp, pos)

	actual := brd.GetOccupiedArea()
	expected := [][]bool{
		{true, false, false, false},
		{true, true, false, false},
		{true, true, true, false},
		{false, false, false, false},
	}

	row, column := brd.GetSize()
	isSuccess := true
	for i := 0; i < row; i++ {
		if isSuccess {
			for j := 0; j < column; j++ {
				isSuccess = actual[i][j] == expected[i][j]
				if !isSuccess {
					break
				}
			}
		}
	}

	if !isSuccess {
		t.Errorf("board should be the same")
	}
}

func TestDoubleBoardUnderMinimimum(t *testing.T) {
	board, _ := board.NewGameBoard(2, 2)
	status := board.DoubleBoard(-1)
	if status {
		t.Errorf("double board should fail (false)")
	}
}

func TestDoubleBoardOverMaximum(t *testing.T) {
	brd, _ := board.NewGameBoard(2, 2)
	status := brd.DoubleBoard(2)
	if status {
		t.Errorf("double board should fail (false)")
	}
}

func TestDobuleBoardRightSuccess(t *testing.T) {
	brd, _ := board.NewGameBoard(2, 2)
	status := brd.DoubleBoard(board.ToRight)
	row, column := brd.GetSize()

	if !(status && row == 2 && column == 4) {
		t.Errorf("expected: status true, row 2, column 4 -"+
			" actual: status %t, row %d, column %d", status, row, column)
	}
}

func TestDobuleBoardRightFail(t *testing.T) {
	brd := board.GameBoard{}
	status := brd.DoubleBoard(board.ToRight)

	if status {
		t.Errorf("expected: false")
	}
}

func TestDobuleBoardButtomSuccess(t *testing.T) {
	brd, _ := board.NewGameBoard(2, 2)
	status := brd.DoubleBoard(board.ToButtom)
	row, column := brd.GetSize()

	if !(status && row == 4 && column == 2) {
		t.Errorf("expected: status true, row 4, column 2 -"+
			" actual: status %t, row %d, column %d", status, row, column)
	}
}

func TestDobuleBoardButtomFail(t *testing.T) {
	brd := board.GameBoard{}
	status := brd.DoubleBoard(board.ToButtom)

	if status {
		t.Errorf("expected: false")
	}
}

func TestAddCalculateOccupiedAreaSuccess(t *testing.T) {
	sp, _ := rectangle.NewRectangle(3, 3)
	brd, _ := board.NewGameBoard(10, 10)
	pos, _ := position.NewPosition(0, 0)

	brd.AddShape(sp, pos)
	area := brd.CalculateOccupiedArea()
	if area != 9 {
		t.Errorf("expected: 9 - actual: %d", area)
	}
}

func TestAddCalculateOccupiedAreaFail(t *testing.T) {
	brd := board.GameBoard{}
	area := brd.CalculateOccupiedArea()
	if area != -1 {
		t.Errorf("expected: -1 - actual: %d", area)
	}
}

func TestGetSizeSuccess(t *testing.T) {
	brd, _ := board.NewGameBoard(10, 10)
	row, column := brd.GetSize()
	if !(row == 10 && column == 10) {
		t.Errorf("expected: (10, 10) - actual: (%d, %d)", row, column)
	}
}

func TestGetSizeFail(t *testing.T) {
	brd := board.GameBoard{}
	row, column := brd.GetSize()
	if !(row == -1 && column == -1) {
		t.Errorf("expected: (-1, -1) - actual: (%d, %d)", row, column)
	}
}
