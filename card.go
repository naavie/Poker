package main

// Card Constructer - Takes a string such as "2H" that returns a card object that returns a number and a suit.
// The number should be comparable to the >, <, ,etc and the suit should have equality.

// Example of Constraint: Given card, 10H, return
//{ Rank:ten, Suit: hearts}

type Card struct {
	rank Rank
	suit Suit
}

// Strings are slices, you can pick out parts by normal slice mechanics... str[2:] (give me everything from the second to the last) or str[1:len(str)] (second to the last).

func CardConstructer(card string) Card {
	//example := "10H"
	// turn the variable example into 10 & H, separately. Output: { Rank:ten, Suit: hearts}

}

type Suit struct {
	suit string
}

var suits = map[string]Suit{
	"H": Suit{"Hearts"},
	"S": Suit{"Spades"},
	"C": Suit{"Clubs"},
	"D": Suit{"Diamonds"},
}

type Rank struct {
	value int
	name  string
}

func (r Rank) LessThan(r2 Rank) bool {
	return r.value < r2.value
}

func (r Rank) GreaterThan(r2 Rank) bool {
	return r.value > r2.value
}

var ranks = map[string]Rank{
	"2":  Rank{2, "Two"},
	"3":  Rank{3, "Three"},
	"4":  Rank{4, "Four"},
	"5":  Rank{5, "Five"},
	"6":  Rank{6, "Six"},
	"7":  Rank{7, "Seven"},
	"8":  Rank{8, "Eight"},
	"9":  Rank{9, "Nine"},
	"10": Rank{10, "Ten"},
	"J":  Rank{11, "Jack"},
	"Q":  Rank{12, "Queen"},
	"K":  Rank{13, "King"},
	"A":  Rank{14, "Ace"},
}
