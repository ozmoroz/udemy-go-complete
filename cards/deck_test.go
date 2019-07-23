package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_desktesting")
	deck := newDeck()
	deck.saveToFile("_desktesting")

	loadedDeck := newDeckFromFile("_desktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 cards, got %d", len(loadedDeck))
	}
	os.Remove("_desktesting")
}
