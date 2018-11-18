package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfyNewDeck(t *testing.T) {
	assert := assert.New(t)
	deck := newDeck()

	assert.Equal(16, len(deck))
	assert.Equal("Ace of Spades", deck[0], "First card in deck is wrong")
	lastCard := deck[len(deck)-1]
	assert.Equal("Four of Clubs", lastCard, "Last card in deck is wrong")
}

func TestIfyWriteToFileAndReadDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.writeToFile(filename)

	loadedDeck := readDeckFromFile(filename)

	assert.Equal(t, len(deck), len(loadedDeck), "Loaded deck should have same length")

	os.Remove(filename)
}
