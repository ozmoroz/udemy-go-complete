package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")
	// fmt.Println(cards.toString())

	cards := newDeckFromFile("my_cards.txt")
	cards.print()

	cards.shuffle()
	fmt.Println("Shuffled deck")
	cards.print()

	// hand.print()
	// remainingCards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
