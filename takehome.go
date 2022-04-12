package main

import (
	"fmt"
	"strings"
)

type Frame struct {
	BlackHand string
	WhiteHand string
}

// Take Home: Using the initial hands string, write a text function, composed of other functions that turns it into two hands that are identifiable into white hand and black hand.

// Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH

func main() {
	BlackHand := []string{"2H", "3D", "5S", "9C", "KD"}
	// BlackHand[0] = "2H"
	// BlackHand[1] = "3D"
	// BlackHand[2] = "5S"
	// BlackHand[3] = "9C"
	// BlackHand[4] = "KD"

	WhiteHand := []string{"2C", "3H", "4S", "8C", "AH"}
	// WhiteHand[0] = "2C"
	// WhiteHand[1] = "3H"
	// WhiteHand[2] = "4S"
	// WhiteHand[3] = "8C"
	// WhiteHand[4] = "AH"

	playerOne := "Black:"
	playerTwo := "White:"

	if len(BlackHand) == 5 { //5 cards
		BlackHand := strings.Join(BlackHand, ", ")
		fmt.Println(playerOne, BlackHand)
	}
	if len(WhiteHand) == 5 { // 5 cards
		WhiteHand := strings.Join(WhiteHand, ", ")
		fmt.Println(playerTwo, WhiteHand)
	}

	if len(BlackHand) == 5 && len(WhiteHand) == 5 {
		BlackHand := strings.Join(BlackHand, ", ")
		WhiteHand := strings.Join(WhiteHand, ", ")
		fmt.Println(playerOne, BlackHand, playerTwo, WhiteHand)
	}

}

// Task One: Separate every element in originalBlackHand & originalWhiteHand with the SPLIT function such that there is a comma between the former and latter elements

// Using the split function one additional time, this is the output you would expect:

//																					   ["Black:","2H","3D","5S","9C","KD","White:","2C","3H","4S","8C","AH"]

// SPLIT function
