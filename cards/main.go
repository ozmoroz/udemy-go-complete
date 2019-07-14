package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
	// fmt.Println(cards.toString())

	// hand.print()
	// remainingCards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
