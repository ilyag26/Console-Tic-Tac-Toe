package main

import (
	a "TicTacToeGo/app"
	"fmt"
	"math/rand"

	"golang.org/x/exp/slices"
)

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
	a.PrintFields(fieldExemple)
	//printing field
	fmt.Println("Field of game:")
	a.PrintFields(fieldGame)
	//printing game
	for !gameIn {
		checkBoolPose = false
		fmt.Println("Make step:")
		fmt.Scan(&step)
		unikCoordinate = append(unikCoordinate, step)
		fmt.Println("Your step:")
		a.MakeStep(step, fieldGame, "X")
		a.PrintFields(fieldGame)
		gameIn = a.CheckWin(gameIn, fieldGame, unikCoordinate)
		if gameIn {
			break
		}

		fmt.Println("Bot's step:")
		for !checkBoolPose {
			randPose = rand.Intn(max-min) + min
			if !slices.Contains(unikCoordinate, randPose) {
				unikCoordinate = append(unikCoordinate, randPose)
				checkBoolPose = true
			}
		}
		a.MakeStep(randPose, fieldGame, "O")
		a.PrintFields(fieldGame)
		gameIn = a.CheckWin(gameIn, fieldGame, unikCoordinate)
		if gameIn {
			break
		}
	}
}
