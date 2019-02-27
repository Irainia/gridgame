package util

import (
	"math/rand"

	"github.com/Irainia/gridgame/shape/circle"

	"github.com/Irainia/gridgame/shape"
	rectangle "github.com/Irainia/gridgame/shape/rectangle"
	triangle "github.com/Irainia/gridgame/shape/triangle"
	"github.com/Irainia/gridgame/util"
)

// GetRandomShape will generate random shape with random size
func GetRandomShape(maxRow, maxColumn int) shape.IShape {
	randomNumber := rand.Intn(util.MaxByte)
	minRow, minColumn := 4, 4
	randRow, randColumn := rand.Intn(maxRow-minRow+1)+minRow,
		rand.Intn(maxColumn-minColumn+1)+minColumn
	switch randomNumber % 6 {
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
	case 5:
		// circle
		cir, _ := circle.NewCircle(randRow)
		return cir
	}

	return nil
}
