package board

import (
	"errors"

	"github.com/irainia/gridgame/position"
	"github.com/irainia/gridgame/shape"
)

// EDoubleDirection will direct doubling the board
type EDoubleDirection int

const (
	// ToRight will direct double to right
	ToRight EDoubleDirection = iota
	// ToButtom will direct double to buttom
	ToButtom
)

// GameBoard will manage board
type GameBoard struct {
	row, column  *int
	gridBoard    [][]bool
	listShape    []shape.IShape
	listPosition []position.Position
}

// NewGameBoard will create Game
func NewGameBoard(row, column int) (*GameBoard, error) {
	if row <= 0 {
		return nil, errors.New("row should be more than zero")
	}
	if column <= 0 {
		return nil, errors.New("column should be more than zero")
	}

	gridBoard := make([][]bool, row)
	for i := 0; i < row; i++ {
		gridBoard[i] = make([]bool, column)
	}

	output := GameBoard{}
	output.row, output.column = &row, &column
	output.gridBoard = gridBoard
	return &output, nil
}

// AddShape will add shape to the board
func (g *GameBoard) AddShape(sp shape.IShape, pos *position.Position) (bool, error) {
	if g.row == nil || g.column == nil {
		return true, errors.New("board is not initialized")
	}
	if sp == nil {
		return true, errors.New("shape should not be a nil value")
	}
	if pos == nil {
		return true, errors.New("position should not be a nil value")
	}

	posCoordinate, err := pos.GetPosition()
	if err != nil {
		return true, errors.New("position is not initialized")
	}

	if posCoordinate[0] >= *g.row || posCoordinate[1] >= *g.column {
		return true, errors.New("position exceed board")
	}

	spRow, spColumn := sp.GetSize()
	if spRow == -1 || spColumn == -1 {
		return true, errors.New("shape is not initialized")
	}

	if posCoordinate[0]+spRow-1 >= *g.row ||
		posCoordinate[1]+spColumn-1 >= *g.column {
		return true, errors.New("position exceed board")
	}

	spOccupied := sp.GetOccupiedArea()
	for i := 0; i < spRow; i++ {
		for j := 0; j < spColumn; j++ {
			if spOccupied[i][j] {
				if g.gridBoard[i+posCoordinate[0]][j+posCoordinate[1]] {
					return false, errors.New("overlap adding shape")
				}
			}
		}
	}

	for i := 0; i < spRow; i++ {
		for j := 0; j < spColumn; j++ {
			if spOccupied[i][j] {
				g.gridBoard[i+posCoordinate[0]][j+posCoordinate[1]] = spOccupied[i][j]
			}
		}
	}

	g.listShape = append(g.listShape, sp)
	g.listPosition = append(g.listPosition, *pos)

	return true, nil
}

// DoubleBoard will double the board based on direction
func (g *GameBoard) DoubleBoard(direct EDoubleDirection) bool {
	if g.row == nil || g.column == nil {
		return false
	}
	var newBoard [][]bool
	var newRow, newColumn int

	switch direct {
	case ToRight:
		newRow = *g.row
		newColumn = *g.column * 2
	case ToButtom:
		newRow = *g.row * 2
		newColumn = *g.column
	default:
		return false
	}

	newBoard = make([][]bool, newRow)
	for i := 0; i < newRow; i++ {
		newBoard[i] = make([]bool, newColumn)
	}

	for i := 0; i < *g.row; i++ {
		for j := 0; j < *g.column; j++ {
			newBoard[i][j] = g.gridBoard[i][j]
		}
	}

	g.gridBoard = newBoard
	g.row = &newRow
	g.column = &newColumn

	return true
}

// CalculateOccupiedArea will return number of area occupied
func (g *GameBoard) CalculateOccupiedArea() int {
	if g.row == nil || g.column == nil {
		return -1
	}

	output := 0
	for i := 0; i < *g.row; i++ {
		for j := 0; j < *g.column; j++ {
			if g.gridBoard[i][j] {
				output++
			}
		}
	}

	return output
}

// GetOccupiedArea will return occupied area on the board
func (g *GameBoard) GetOccupiedArea() [][]bool {
	if g.row == nil || g.column == nil {
		return nil
	}

	output := make([][]bool, *g.row)
	for i := 0; i < *g.row; i++ {
		output[i] = make([]bool, *g.column)
	}
	copy(output, g.gridBoard)
	return output
}

// GetSize will return board size (row, column)
func (g *GameBoard) GetSize() (int, int) {
	if g.row == nil || g.column == nil {
		return -1, -1
	}
	return *g.row, *g.column
}
