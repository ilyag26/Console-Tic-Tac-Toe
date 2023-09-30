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

func MakeStepUser(step int, field [][]string) {
	switch {
	case step == 1:
		field[0][0] = "X"
	case step == 2:
		field[0][1] = "X"
	case step == 3:
		field[0][2] = "X"
	case step == 4:
		field[1][0] = "X"
	case step == 5:
		field[1][1] = "X"
	case step == 6:
		field[1][2] = "X"
	case step == 7:
		field[2][0] = "X"
	case step == 8:
		field[2][1] = "X"
	case step == 9:
		field[2][2] = "X"
	}
}

func MakeStepBot(randomNumber int, field [][]string) {
	switch {
	case randomNumber == 1:
		field[0][0] = "O"
	case randomNumber == 2:
		field[0][1] = "O"
	case randomNumber == 3:
		field[0][2] = "O"
	case randomNumber == 4:
		field[1][0] = "O"
	case randomNumber == 5:
		field[1][1] = "O"
	case randomNumber == 6:
		field[1][2] = "O"
	case randomNumber == 7:
		field[2][0] = "O"
	case randomNumber == 8:
		field[2][1] = "O"
	case randomNumber == 9:
		field[2][2] = "O"
	}
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
