package util

import (
	"math/rand"

	"github.com/irainia/gridgame/shape"
	rectangle "github.com/irainia/gridgame/shape/rectangle"
	triangle "github.com/irainia/gridgame/shape/triangle"
	"github.com/irainia/gridgame/util"
)

// GetRandomShape will generate random shape with random size
func GetRandomShape(maxRow, maxColumn int) shape.IShape {
	randomNumber := rand.Intn(util.MaxByte)
	minRow, minColumn := 2, 2
	randRow, randColumn := rand.Intn(maxRow-minRow+1)+minRow,
		rand.Intn(maxColumn-minColumn+1)+minColumn
	switch randomNumber % 5 {
	case 0:
		// rectangle
		rect, _ := rectangle.NewRectangle(randRow, randColumn)
		return rect
	case 1:
		// left buttom triangle
		tri, _ := triangle.NewLeftButtomTriangle(randRow, randColumn)
		return tri
	case 2:
		// left top triangle
		tri, _ := triangle.NewLeftTopTriangle(randRow, randColumn)
		return tri
	case 3:
		// right buttom triangle
		tri, _ := triangle.NewRightButtomTriangle(randRow, randColumn)
		return tri
	case 4:
		// right top triangle
		tri, _ := triangle.NewRightTopTriangle(randRow, randColumn)
		return tri
	}

	return nil
}
