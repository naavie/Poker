package main

import (
	"fmt"
	"strings"
)

func main() {

	initialHandString := []string{"Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH"}
	fmt.Println(initialHandString)

	fmt.Println("The length of the slice, initialHandString is:", len(initialHandString))

	// Turning a slice of length 1 into a slice of length 12 by separating every element by a comma

	for i, ch := range initialHandString {
		strings.Join(initialHandString, ", ")
		fmt.Println(ch, i)

	}

}

func MakeNewCard(number string, suit string) {

	deckcardNumber := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	deckcardSuit := []string{"H", "D", "S", "C"}

	fmt.Println(deckcardNumber)
	fmt.Println(deckcardSuit)

}
