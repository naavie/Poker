package main

import "fmt"

func main() {
	s := make([]string, 10)

	s[0] = "5H"
	s[1] = "2C"
	s[2] = "10D"
	s[3] = "JC"
	s[4] = "7H"
	s[5] = "10H"
	s[6] = "KC"
	s[7] = "QH"
	s[8] = "9C"
	s[9] = "2D"

	if len(s)%5 == 0 {
		fmt.Println("This code works")

		playerName := "Navmeet: "
		fmt.Println(playerName, s[:5])
		if len(s)%5 != 5 {
			fmt.Println("This code works too! - TESTING")
			playerName := "Sean: "
			fmt.Println(playerName, s[5:])
		}
	}
}

// type HandSeperation struct {
// 	playerName string
// 	initalHand []string
// }

// func HandSeperator(playerName string, initalHand []string) HandSeperation {
// 	if len(initalHand)%5 == 0 {
// 		playerName = "Player One: "
// 		playerOneCards := initalHand[:5]
// 		return HandSeperation{playerName: playerName, initalHand: playerOneCards}
// 	} else if len(initalHand)%5 != 0 {
// 		playerName = "Player Two: "
// 		playerTwoCards := initalHand[5:]
// 		return HandSeperation{playerName: playerName, initalHand: playerTwoCards}
// 	}
// }
//---------------------------------

//

// type Seperator struct {
// 	playerOneName  string
// 	playerTwoName  string
// 	playerOneCards []Card
// 	playerTwoCards []Card
// }

// func HandSeperator(p1Name string, p2Name string, initalHand []Card) Seperator {

// 	p1Cards := initalHand[:5]
// 	p2Cards := initalHand[5:]

// 	return Seperator{playerOneName: p1Name, playerOneCards: p1Cards, playerTwoName: p2Name, playerTwoCards: p2Cards}
// }
