package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Error("Expected deck of length 12 bet got len", len(d))
	}

	if d[0] != "ace of spades" {
		t.Error("Expected first card to be ace of spades but got", d[0])
	}

	if d[len(d)-1] != "four of hearts" {
		t.Error("Expected last card to be four of hearts but got", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Error("Expected deck of length 12 bet got len", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
