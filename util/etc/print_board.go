package util

import "fmt"

// PrintBoard will print board
func PrintBoard(brd [][]bool) {
	fmt.Print("  ")
	if len(brd) > 9 {
		fmt.Print(" ")
	}

	for i := 0; i < len(brd[0]); i++ {
		fmt.Printf("%d ", i+1)
		if len(brd[0]) > 9 && i < 9 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()

	for i := 0; i < len(brd); i++ {
		fmt.Printf("%d ", i+1)
		if len(brd) > 9 && i < 9 {
			fmt.Printf(" ")
		}

		for j := 0; j < len(brd[i]); j++ {
			if brd[i][j] {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}

			if len(brd[0]) > 9 && j < 9 {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}
