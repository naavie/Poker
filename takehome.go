package main

import (
	"fmt"
	"strings"
)

// Take Home: Using the initial hands string, write a text function, composed of other functions that turns it into two hands that are identifiable into white hand and black hand.

// Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH

func main() {
	//  Variables that contain the original card set we were working with.
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

	// The print statements below should match the expected output indicated in line 7.

	playerOne := "Black:"
	playerTwo := "White:"

	// Testing Pt. 1: Code works this far.
	fmt.Println(playerOne, BlackHand)
	fmt.Println(playerTwo, WhiteHand)

	// Testing Pt. 1: Continued.
	fmt.Println(playerOne, BlackHand, playerTwo, WhiteHand)

	BlackHandStrings := strings.Join(BlackHand, ", ")
	WhiteHandStrings := strings.Join(WhiteHand, ", ")

	fmt.Println(BlackHandStrings)
	fmt.Println(WhiteHandStrings)

	// Final steps

}

// Task One: Separate every element in originalBlackHand & originalWhiteHand with the SPLIT function such that there is a comma between the former and latter elements

// Using the split function one additional time, this is the output you would expect:

//																					   ["Black:","2H","3D","5S","9C","KD","White:","2C","3H","4S","8C","AH"]

// SPLIT function

func Split(string) {

}
