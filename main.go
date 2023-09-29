package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func printFields(fields [][]string) {
	for i := 0; i < len(fields); i++ {
		for j := 0; j < len(fields); j++ {
			fmt.Print(fields[i][j] + " ")
		}
		fmt.Print("\n")
	}
}

func makeStepUser(step int, field [][]string) {
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

func makeStepBot(randomNumber int, field [][]string) {
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

func checkWin(gameIn1 bool, field [][]string, unikCoordinate []int) bool {
	if field[0][0] == "X" && field[0][1] == "X" && field[0][2] == "X" {
		PrintWinX()
		return true
	} else if field[0][0] == "O" && field[0][1] == "O" && field[0][2] == "O" {
		PrintWinO()
		return true
	} else if field[1][0] == "X" && field[1][1] == "X" && field[1][2] == "X" {
		PrintWinX()
		return true
	} else if field[1][0] == "O" && field[1][1] == "O" && field[1][2] == "O" {
		PrintWinO()
		return true
	} else if field[2][0] == "X" && field[2][1] == "X" && field[2][2] == "X" {
		PrintWinX()
		return true
	} else if field[2][0] == "O" && field[2][1] == "O" && field[2][2] == "O" {
		PrintWinO()
		return true
	} else if field[0][0] == "X" && field[1][1] == "X" && field[2][2] == "X" {
		PrintWinX()
		return true
	} else if field[0][0] == "O" && field[1][1] == "O" && field[2][2] == "O" {
		PrintWinO()
		return true
	} else if field[0][2] == "X" && field[1][1] == "X" && field[2][0] == "X" {
		PrintWinX()
		return true
	} else if field[0][2] == "O" && field[1][1] == "O" && field[2][0] == "O" {
		PrintWinO()
		return true
	} else if field[0][0] == "X" && field[1][0] == "X" && field[2][0] == "X" {
		PrintWinX()
		return true
	} else if field[0][0] == "O" && field[1][0] == "O" && field[2][0] == "O" {
		PrintWinO()
		return true
	} else if field[0][2] == "X" && field[1][2] == "X" && field[2][2] == "X" {
		PrintWinX()
		return true
	} else if field[0][2] == "O" && field[1][2] == "O" && field[2][2] == "O" {
		PrintWinO()
		return true
	} else if len(unikCoordinate) == 9 {
		PrintDraw()
		return true
	}
	return false
}

func PrintWinX() {
	fmt.Println("X <- WINNER")
}

func PrintWinO() {
	fmt.Println("O <- WINNER")
}

func PrintDraw() {
	fmt.Println("GAME END! DRAW")
}

func main() {
	//for detecting step
	var step int
	// field for game
	fieldGame := [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	// field for rules
	fieldExemple := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	//unik coordinte
	var unikCoordinate []int
	//min random number
	min := 1
	//max random number
	max := 9
	//game check
	gameIn := false
	//random coordinates for position of bot's steps
	randPose := 0
	//check unik bool
	var checkBoolPose bool
	//printing name
	fmt.Println("-----GO TIC-TAC-TOE-----")
	//printing rules
	fmt.Println("For make step just enter folowing numbers for each position:")
	printFields(fieldExemple)
	//printing field
	fmt.Println("Field of game:")
	printFields(fieldGame)
	//printing game
	for gameIn == false {
		checkBoolPose = false
		fmt.Println("Make step:")
		fmt.Scan(&step)
		unikCoordinate = append(unikCoordinate, step)
		fmt.Println("Your step:")
		makeStepUser(step, fieldGame)
		printFields(fieldGame)
		gameIn = checkWin(gameIn, fieldGame, unikCoordinate)
		if gameIn {
			break
		}

		fmt.Println("Bot's step:")
		for !checkBoolPose {
			randPose = rand.Intn(max-min) + min
			if slices.Contains(unikCoordinate, randPose) == false {
				unikCoordinate = append(unikCoordinate, randPose)
				checkBoolPose = true
			}
		}
		makeStepBot(randPose, fieldGame)
		printFields(fieldGame)
		gameIn = checkWin(gameIn, fieldGame, unikCoordinate)
		if gameIn {
			break
		}
	}
}
