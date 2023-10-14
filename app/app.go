package app

import (
	"fmt"
)

func PrintFields(fields [][]string) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields); j++ {
			fmt.Print(fields[i][j] + " ")
		}
		fmt.Print("\n")
	}
}

func MakeStep(step int, field [][]string, symbol string) {

	m := map[int][2]int{
		1: {0, 0},
		2: {0, 1},
		3: {0, 2},
		4: {1, 0},
		5: {1, 1},
		6: {1, 2},
		7: {2, 0},
		8: {2, 1},
		9: {2, 2},
	}

	field[m[step][0]][m[step][1]] = symbol
}

func CheckWin(gameIn1 bool, field [][]string, unikCoordinate []int) bool {
	if field[0][0] == "X" && field[0][1] == "X" && field[0][2] == "X" {
		printWinX()
		return true
	} else if field[0][0] == "O" && field[0][1] == "O" && field[0][2] == "O" {
		printWinO()
		return true
	} else if field[1][0] == "X" && field[1][1] == "X" && field[1][2] == "X" {
		printWinX()
		return true
	} else if field[1][0] == "O" && field[1][1] == "O" && field[1][2] == "O" {
		printWinO()
		return true
	} else if field[2][0] == "X" && field[2][1] == "X" && field[2][2] == "X" {
		printWinX()
		return true
	} else if field[2][0] == "O" && field[2][1] == "O" && field[2][2] == "O" {
		printWinO()
		return true
	} else if field[0][0] == "X" && field[1][1] == "X" && field[2][2] == "X" {
		printWinX()
		return true
	} else if field[0][0] == "O" && field[1][1] == "O" && field[2][2] == "O" {
		printWinO()
		return true
	} else if field[0][2] == "X" && field[1][1] == "X" && field[2][0] == "X" {
		printWinX()
		return true
	} else if field[0][2] == "O" && field[1][1] == "O" && field[2][0] == "O" {
		printWinO()
		return true
	} else if field[0][0] == "X" && field[1][0] == "X" && field[2][0] == "X" {
		printWinX()
		return true
	} else if field[0][0] == "O" && field[1][0] == "O" && field[2][0] == "O" {
		printWinO()
		return true
	} else if field[0][2] == "X" && field[1][2] == "X" && field[2][2] == "X" {
		printWinX()
		return true
	} else if field[0][2] == "O" && field[1][2] == "O" && field[2][2] == "O" {
		printWinO()
		return true
	} else if len(unikCoordinate) == 9 {
		printDraw()
		return true
	}
	return false
}

func printWinX() {
	fmt.Println("X <- WINNER")
}

func printWinO() {
	fmt.Println("O <- WINNER")
}

func printDraw() {
	fmt.Println("GAME END! DRAW")
}
