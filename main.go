package main

import (
	"fmt"
	"math/rand"
)

var playerOneHand string = "Black:"
var playerTwoHand string = "White:"

// Example: A hand is represented by a slice of cards, size 5
// -----
//			Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH

// Function: Turing a string into a card: How do I turn the string into a card?

//- The first character of a card string is its number and the second character is the card’s suit.

// - To convert a string into a card, we can create a new card and select the 0th and 1st elements.
// - The 0th element can be used to set the card’s number and the 1st element can be used to set the card’s suit.

func newCard() {

	cardNumber := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	cardSuit := []string{"H", "D", "S", "C"}

	for i := 0; i <= 5; i++ {
		randomCardIndex := rand.Intn(len(cardNumber))
		randomSuitIndex := rand.Intn(len(cardSuit))

		playerCardNumber := cardNumber[randomCardIndex]
		playerCardSuit := cardSuit[randomSuitIndex]

		playerHand := playerCardNumber + playerCardSuit

	} playerHand

}
