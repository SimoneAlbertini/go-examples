package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	firstCard := deck[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first element 'Ace of Spades', but got '%v'", firstCard)
	}

	lastCard := deck[len(deck)-1]
	if lastCard != "Four of Clubs" {
		t.Errorf("Expected deck length of 16, but got %v", lastCard)
	}
}

func TestWriteToFileAndReadDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.writeToFile(filename)

	loadedDeck := readDeckFromFile(filename)

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected loaded deck to have length %v, but got %v", len(loadedDeck), len(deck))
	}

	os.Remove(filename)
}
