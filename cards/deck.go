package main

import "fmt"

// Create a new type of 'deck' - a slice of strings
type deck []string

// create a new deck
func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// print deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
