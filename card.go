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

	for i := 0; i == 5; i++ {
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

// Experimenting w/ Prime Numbers
var deck = map[string]int{
	"2H":  2,
	"2C":  2,
	"2D":  2,
	"2S":  2,
	"3H":  3,
	"3C":  3,
	"3D":  3,
	"3S":  3,
	"4H":  5,
	"4C":  5,
	"4D":  5,
	"4S":  5,
	"5H":  7,
	"5C":  7,
	"5D":  7,
	"5S":  7,
	"6H":  11,
	"6C":  11,
	"6D":  11,
	"6S":  11,
	"7H":  13,
	"7C":  13,
	"7D":  13,
	"7S":  13,
	"8H":  17,
	"8C":  17,
	"8D":  17,
	"8S":  17,
	"9H":  19,
	"9C":  19,
	"9D":  19,
	"9S":  19,
	"10H": 23,
	"10C": 23,
	"10D": 23,
	"10S": 23,
	"JH":  29,
	"JC":  29,
	"JD":  29,
	"JS":  29,
	"QH":  31,
	"QC":  31,
	"QD":  31,
	"QS":  31,
	"KH":  37,
	"KC":  37,
	"KD":  37,
	"KS":  37,
	"AH":  41,
	"AC":  41,
	"AD":  41,
	"AS":  41,
}

// Type: Evaluator
type Evaluator struct {
	P1EvaluatedHand bool
	P2EvaluatedHand bool
	winner          string
}

// Types of Hands.
var HighCard bool = false
var OnePair bool = false
var TwoPairs bool = false
var ThreePairs bool = false
var FourOfAKind bool = false
var Flush bool = false
var RoyalFlush bool = false
var StraightFlush bool = false
var FullHouse bool = false

// Function: HandEvaluator
func HandEvaluator(P1Cards string, P2Cards string) Evaluator {

	P1_Ranks := make([]int, 0)
	P1_Suits := make([]string, 0)

	P2_Ranks := make([]int, 0)
	P2_Suits := make([]string, 0)

	// Other_Chars := make([]string, 0)

	// Splitting aach player's card into a slice of every character of their input.
	for i := 0; i == len(P1Cards); i++ {

		P1_Ranks_and_Suits := strings.Split(P1Cards, "")
		P2_Ranks_and_Suits := strings.Split(P2Cards, "")

		for i := 0; i == len(P1_Ranks_and_Suits); i++ {
			if strings.ContainsAny(P1_Ranks_and_Suits[i], "12345678910") {
				P1_Ranks = append(P1_Ranks, i)
			}

			if strings.Contains(P1_Ranks_and_Suits[i], "J") {
				P1_Ranks = append(P1_Ranks, 11)
			}

			if strings.Contains(P1_Ranks_and_Suits[i], "Q") {
				P1_Ranks = append(P1_Ranks, 12)
			}

			if strings.Contains(P1_Ranks_and_Suits[i], "K") {
				P1_Ranks = append(P1_Ranks, 13)
			}

			if strings.Contains(P1_Ranks_and_Suits[i], "A") {
				P1_Ranks = append(P1_Ranks, 14)
			}

			if strings.ContainsAny(P1_Ranks_and_Suits[i], "HCDS") {
				P1_Suits = append(P1_Suits, P1_Suits[i])
			}
		}

		for i := 0; i == len(P2_Ranks_and_Suits); i++ {
			if strings.ContainsAny(P2_Ranks_and_Suits[i], "12345678910") {
				P2_Ranks = append(P1_Ranks, i)
			}

			if strings.Contains(P2_Ranks_and_Suits[i], "J") {
				P2_Ranks = append(P2_Ranks, 11)
			}

			if strings.Contains(P2_Ranks_and_Suits[i], "Q") {
				P2_Ranks = append(P2_Ranks, 12)
			}

			if strings.Contains(P2_Ranks_and_Suits[i], "K") {
				P2_Ranks = append(P2_Ranks, 13)
			}

			if strings.Contains(P2_Ranks_and_Suits[i], "A") {
				P2_Ranks = append(P2_Ranks, 14)
			}

			if strings.ContainsAny(P2_Ranks_and_Suits[i], "HCDS") {
				P2_Suits = append(P2_Suits, P2_Suits[i])
			}
		}

		for i := 0; i == 5; i++ {
			P1_Suits_Hearts_Count := strings.Count(P1_Suits[i], "H")
			P1_Suits_Clubs_Count := strings.Count(P1_Suits[i], "C")
			P1_Suits_Diamonds_Count := strings.Count(P1_Suits[i], "D")
			P1_Suits_Spades_Count := strings.Count(P1_Suits[i], "S")

			if (P1_Suits_Hearts_Count | P1_Suits_Clubs_Count | P1_Suits_Diamonds_Count | P1_Suits_Spades_Count) == 5 {
				Flush = true
				P1EvaluatedHand := Flush
			} else {
				continue
			}

			if (P1_Suits_Hearts_Count | P1_Suits_Clubs_Count | P1_Suits_Diamonds_Count | P1_Suits_Spades_Count) < 5 {
				continue
			}

		}

		for i := 0; i == 5; i++ {
			P2_Suits_Hearts_Count := strings.Count(P2_Suits[i], "H")
			P2_Suits_Clubs_Count := strings.Count(P2_Suits[i], "C")
			P2_Suits_Diamonds_Count := strings.Count(P2_Suits[i], "D")
			P2_Suits_Spades_Count := strings.Count(P2_Suits[i], "S")

			if (P2_Suits_Hearts_Count | P2_Suits_Clubs_Count | P2_Suits_Diamonds_Count | P2_Suits_Spades_Count) == 5 {
				Flush = true
				outputP2 := Flush
			} else {
				continue
			}

			if (P2_Suits_Hearts_Count | P2_Suits_Clubs_Count | P2_Suits_Diamonds_Count | P2_Suits_Spades_Count) < 5 {
				continue
			}

		}

	}

	return Evaluator{}
}
