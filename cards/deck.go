package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// deal the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save deck to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
