package main

import "fmt"

// Take Home: Using the initial hands string, write a text function, composed of other functions that turns it into two hands that are identifiable into white hand and black hand.

// Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH

func main() {
	//  Variables that contain the original card set we were working with.
	var originalBlackHand [5]string

	originalBlackHand[0] = "2H"
	originalBlackHand[1] = "3D"
	originalBlackHand[2] = "5S"
	originalBlackHand[3] = "9C"
	originalBlackHand[4] = "KD"

	var originalWhiteHand [5]string

	originalWhiteHand[0] = "2C"
	originalWhiteHand[1] = "3H"
	originalWhiteHand[2] = "4S"
	originalWhiteHand[3] = "8C"
	originalWhiteHand[4] = "AH"

	fmt.Println(originalBlackHand)
	fmt.Println(originalWhiteHand)
}
