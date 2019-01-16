package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/irainia/gridgame/board"
	"github.com/irainia/gridgame/position"
	"github.com/irainia/gridgame/shape"
	"github.com/irainia/gridgame/util"
	etcUtil "github.com/irainia/gridgame/util/etc"
	ioUtil "github.com/irainia/gridgame/util/io"
)

type concmessage struct {
	Message string
	Shape   shape.IShape
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scoreFile := "log.save"

	maxScore := -1
	score, err := ioUtil.LoadScore(scoreFile)
	if err == nil {
		maxScore = score
		fmt.Printf("Highest score is: %d\n", maxScore)
	}

	cpuBoard, err := board.NewGameBoard(util.BaseBoardRow, util.BaseBoardColumn)
	if err != nil {
		log.Fatal(err)
	}
	playerBoard, err := board.NewGameBoard(util.BaseBoardRow, util.BaseBoardColumn)
	if err != nil {
		log.Fatal(err)
	}

	numberOfPreviousShapesRegistered := 0

	var listCPUShape []shape.IShape
	minRatioOccupied := 0.85

	// game until player lose
	for {
		// create cpuBoard reference
		boardRow, boardColumn := cpuBoard.GetSize()
		maxAddTrial := util.MaxByte

		// trying add
		for {
			maxRandomTrial := util.MaxByte
			var uniqueShape shape.IShape

			// trying random
			for {
				sp := etcUtil.GetRandomShape(boardRow, boardColumn)
				spOccupied := sp.GetOccupiedArea()
				spRow, spColumn := sp.GetSize()

				// get unique shape
				isExist := true
				lenUnique := len(listCPUShape)
				for i := 0; i < lenUnique; i++ {
					if !isExist {
						break
					}

					tp := listCPUShape[i]
					tpOccupied := tp.GetOccupiedArea()
					tpRow, tpColumn := tp.GetSize()

					if spRow != tpRow || spColumn != tpColumn {
						isExist = false
						break
					}

					for p := 0; p < spRow; p++ {
						if isExist {
							for q := 0; q < spColumn; q++ {
								if spOccupied[p][q] != tpOccupied[p][q] {
									isExist = false
								}
							}
						} else {
							break
						}
					}
				}
				maxRandomTrial--
				if isExist && maxRandomTrial > 0 {
					continue
				}
				uniqueShape = sp
				break
			}

			// add unique shape
			isAdded := false
			for p := 0; p < boardRow; p++ {
				if !isAdded {
					for q := 0; q < boardColumn; q++ {
						pos, _ := position.NewPosition(p, q)
						addstatus, err := cpuBoard.AddShape(uniqueShape, pos)
						if addstatus && err == nil {
							listCPUShape = append(listCPUShape, uniqueShape)
							isAdded = true
							break
						}
					}
				}
			}

			if isAdded {
				continue
			}

			maxAddTrial--
			if maxAddTrial < 0 {
				break
			}

			temporaryNumber := float64(cpuBoard.CalculateOccupiedArea())
			if temporaryNumber/float64(boardRow*boardColumn) > minRatioOccupied {
				break
			}
		}

		// begin playerBoard
		fmt.Println("Here's your main board:")
		etcUtil.PrintBoard(playerBoard.GetOccupiedArea())

		listPlayerShape := listCPUShape[numberOfPreviousShapesRegistered:]
		numberOfPreviousShapesRegistered = len(listCPUShape)

		doesPlayerWinStage := true
		var input string
		for i := 0; i < len(listPlayerShape); i++ {
			fmt.Printf("Shape %d of %d:\n", i+1, len(listPlayerShape))
			occupied := listPlayerShape[i].GetOccupiedArea()
			etcUtil.PrintBoard(occupied)

			fmt.Print("Coordinate in row,column: ")
			fmt.Scanln(&input)
			separatedInput := strings.Split(input, ",")
			if len(separatedInput) != 2 {
				fmt.Println("error format input")
				i--
				continue
			}

			row, errRow := strconv.Atoi(separatedInput[0])
			column, errColumn := strconv.Atoi(separatedInput[1])
			row--
			column--
			if !(errRow == nil && errColumn == nil) {
				fmt.Println("error format input")
				i--
				continue
			}
			pos, errPos := position.NewPosition(row, column)
			if errPos != nil {
				fmt.Println("coordinate should be within 1-maxColumn and 1-maxRow")
				i--
				continue
			}

			status, errAdd := playerBoard.AddShape(listPlayerShape[i], pos)
			if errAdd != nil && status {
				fmt.Printf("error adding: %s\n", errAdd.Error())
				i--
				continue
			}

			if !status {
				doesPlayerWinStage = false
				fmt.Println("Shape overlapped, you lose!")
				break
			}

			fmt.Println("Shape is added.")
			fmt.Println("Your new board:")
			etcUtil.PrintBoard(playerBoard.GetOccupiedArea())
		}

		if !doesPlayerWinStage {
			break
		}

		fmt.Println()
		fmt.Println("Bord is expanding...")
		fmt.Print("Press 'Enter' to continue! ")
		fmt.Scanln()

		// Expand board
		if boardRow > boardColumn {
			cpuBoard.DoubleBoard(board.ToRight)
			playerBoard.DoubleBoard(board.ToRight)
		} else if boardColumn > boardRow {
			cpuBoard.DoubleBoard(board.ToButtom)
			playerBoard.DoubleBoard(board.ToButtom)
		} else {
			if rand.Intn(util.MaxByte)%2 == 0 {
				cpuBoard.DoubleBoard(board.ToRight)
				playerBoard.DoubleBoard(board.ToRight)
			} else {
				cpuBoard.DoubleBoard(board.ToButtom)
				playerBoard.DoubleBoard(board.ToButtom)
			}
		}
	}

	currentBestScore := playerBoard.CalculateOccupiedArea()
	fmt.Println()
	fmt.Printf("Your score is: %d\n", currentBestScore)
	if currentBestScore > maxScore {
		maxScore = currentBestScore
		fmt.Println("This is new record!")

		err = ioUtil.SaveScore(scoreFile, maxScore)
		if err != nil {
			fmt.Printf("cannot save score: %s\n", err.Error())
		}
	}
}
