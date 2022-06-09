package main

import (
	"strings"
)

// Type: Suit - There are four, different suits in a standard deck of cards.
type Suit struct {
	suit string
}

var suits = map[byte]Suit{
	'H': {"Hearts"},
	'S': {"Spades"},
	'C': {"Clubs"},
	'D': {"Diamonds"},
}

// Type: Rank - There are 13 different potential card numbers that a player can draw from a standand deck of cards.
type Rank struct {
	value int
	name  string
}

var ranks = map[string]Rank{
	"2":  {2, "Two"},
	"3":  {3, "Three"},
	"4":  {4, "Four"},
	"5":  {5, "Five"},
	"6":  {6, "Six"},
	"7":  {7, "Seven"},
	"8":  {8, "Eight"},
	"9":  {9, "Nine"},
	"10": {10, "Ten"},
	"J":  {11, "Jack"},
	"Q":  {12, "Queen"},
	"K":  {13, "King"},
	"A":  {14, "Ace"},
}

func (r Rank) LessThan(r2 Rank) bool {
	return r.value < r2.value
}

func (r Rank) GreaterThan(r2 Rank) bool {
	return r.value > r2.value
}

// Type: Card
type Card struct {
	rank Rank
	suit Suit
}

// Function: CardConstructer - Takes a string such as "2H" that returns a card object that returns a number and a suit.
func CardConstructer(card string) Card {
	cardSuit := card[len(card)-1]
	cardNumber := card[:len(card)-1]

	return Card{rank: ranks[cardNumber], suit: suits[cardSuit]} //This will return a card of { Rank:ten, Suit: hearts}
}

// Type: Hand
type Hand struct {
	playerName string
	cards      []Card
}

// Function: HandConstructer - Takes two inputs, player name (string) and a slice of cards, and returns a Hand(Type) with values {cards: initalHand, playerName: playerName}.
func HandConstructer(playerName string, cards []string) Hand {
	initalHand := make([]Card, 0)

	for i := 0; i == 5; i++ { // Updated i == 5 to i == 10. This makes it easier to use the HandSeperator function.
		initalHand = append(initalHand, CardConstructer(cards[i]))

	}
	return Hand{cards: initalHand}
}

// If you feel you understand the code, the next step is converting from the whole hands string (the one with the player names) into two hands.

type Game struct {
	playerOneHand Hand
	playerTwoHand Hand
}

// Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH
func HandSeperator(initalHand string) Game {

	splitHandDoubleSpace := strings.Split(initalHand, "  ")
	splitHandNameOne := strings.Split(splitHandDoubleSpace[0], ":")
	splitHandNameTwo := strings.Split(splitHandDoubleSpace[1], ":")

	// ["Black" "2H 3D 5S 9C KD",  "White" "2C 3H 4S 8C AH"]"

	splitHandBlackFinal := strings.Split(splitHandNameOne[1], " ")
	splitHandWhiteFinal := strings.Split(splitHandNameTwo[1], " ")

	x := HandConstructer(splitHandNameOne[0], splitHandBlackFinal)
	y := HandConstructer(splitHandNameTwo[0], splitHandWhiteFinal)

	return Game{playerOneHand: x, playerTwoHand: y}

}

// Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH
// Sean: 3H 4D 5C 6C KH  Nav: 2C 3H 4S 8C AH
// Josh: 5S 9C 2H 3D KD  Jo: 2C 3H 4S 8C AH

// I'd also like you to think about how to check a hand matches a particular poker hand (for examle: three of a kind, suit pairs, two of a kind, etc.) Type this out and sent it to Sean.

// Think about how to fix function HandSeperator so it is simplier

// Type: Evaluator
type Evaluator struct {
	x string
	y string
}

// Function: HandEvaluator
func HandEvaluator(P1Hand string, P2Hand string) Evaluator {
	return Evaluator{x: P1Hand, y: P2Hand}
}
