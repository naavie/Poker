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

type Deck struct {
	cards []Card
}

var deck = map[string]Rank{
	"2H":  {1, "Two of Hearts"},
	"2C":  {2, "Two of Clubs"},
	"2D":  {3, "Two of Diamonds"},
	"2S":  {4, "Two of Spades"},
	"3H":  {5, "Three of Hearts"},
	"3C":  {6, "Three of Clubs"},
	"3D":  {7, "Three of Diamonds"},
	"3S":  {8, "Three of Spades"},
	"4H":  {9, "Four of Hearts"},
	"4C":  {10, "Four of Clubs"},
	"4D":  {11, "Four of Diamonds"},
	"4S":  {12, "Four of Spades"},
	"5H":  {13, "Five of Hearts"},
	"5C":  {14, "Five of Clubs"},
	"5D":  {15, "Five of Diamonds"},
	"5S":  {16, "Five of Spades"},
	"6H":  {17, "Six of Hearts"},
	"6C":  {18, "Six of Clubs"},
	"6D":  {19, "Six of Diamonds"},
	"6S":  {20, "Six of Spades"},
	"7H":  {21, "Seven of Hearts"},
	"7C":  {22, "Seven of Clubs"},
	"7D":  {23, "Seven of Diamonds"},
	"7S":  {24, "Seven of Spades"},
	"8H":  {25, "Eight of Hearts"},
	"8C":  {26, "Eight of Clubs"},
	"8D":  {27, "Eight of Diamonds"},
	"8S":  {28, "Eight of Spades"},
	"9H":  {29, "Nine of Hearts"},
	"9C":  {30, "Nine of Clubs"},
	"9D":  {31, "Nine of Diamonds"},
	"9S":  {32, "Nine of Spades"},
	"10H": {33, "Ten of Hearts"},
	"10C": {34, "Ten of Clubs"},
	"10D": {35, "Ten of Diamonds"},
	"10S": {36, "Ten of Spades"},
	"JH":  {37, "Jack of Hearts"},
	"JC":  {38, "Jack of Clubs"},
	"JD":  {39, "Jack of Diamonds"},
	"JS":  {40, "Jack of Spades"},
	"QH":  {41, "Queen of Hearts"},
	"QC":  {42, "Queen of Clubs"},
	"QD":  {43, "Queen of Diamonds"},
	"QS":  {44, "Queen of Spades"},
	"KH":  {45, "King of Hearts"},
	"KC":  {46, "King of Clubs"},
	"KD":  {47, "King of Diamonds"},
	"KS":  {48, "King of Spades"},
	"AH":  {49, "Ace of Hearts"},
	"AC":  {50, "Ace of Clubs"},
	"AD":  {51, "Ace of Diamonds"},
	"AS":  {52, "Ace of Spades"},
}

// var ranks = map[string]Rank{
// 	"2":  {2, "Two"},
// 	"3":  {3, "Three"},
// 	"4":  {4, "Four"},
// 	"5":  {5, "Five"},
// 	"6":  {6, "Six"},
// 	"7":  {7, "Seven"},
// 	"8":  {8, "Eight"},
// 	"9":  {9, "Nine"},
// 	"10": {10, "Ten"},
// 	"J":  {11, "Jack"},
// 	"Q":  {12, "Queen"},
// 	"K":  {13, "King"},
// 	"A":  {14, "Ace"},
// }

// Type: Evaluator
type Evaluator struct {
	x string
	y string
}

// Function: HandEvaluator
func HandEvaluator(P1Hand string, P2Hand string) Evaluator {

	var suits_enumerated [4]int
	suits_enumerated[0] = 1
	suits_enumerated[1] = 2
	suits_enumerated[2] = 3
	suits_enumerated[3] = 4

	return Evaluator{}
}

// https://stackoverflow.com/questions/42380183/algorithm-to-give-a-value-to-a-5-card-poker-hand

// type Suit struct {
// 	suit string
// }

// var suits = map[byte]Suit{
// 	'H': {"Hearts"},
// 	'S': {"Spades"},
// 	'C': {"Clubs"},
// 	'D': {"Diamonds"},
// }

// // Type: Rank - There are 13 different potential card numbers that a player can draw from a standand deck of cards.
// type Rank struct {
// 	value int
// 	name  string
// }

// var ranks = map[string]Rank{
// 	"2":  {2, "Two"},
// 	"3":  {3, "Three"},
// 	"4":  {4, "Four"},
// 	"5":  {5, "Five"},
// 	"6":  {6, "Six"},
// 	"7":  {7, "Seven"},
// 	"8":  {8, "Eight"},
// 	"9":  {9, "Nine"},
// 	"10": {10, "Ten"},
// 	"J":  {11, "Jack"},
// 	"Q":  {12, "Queen"},
// 	"K":  {13, "King"},
// 	"A":  {14, "Ace"},
// }
