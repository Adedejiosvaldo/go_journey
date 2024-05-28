package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		// t.Error("Expected deck lenght of 16 , but got", len(d))
		t.Errorf("Expected deck lenght of 16 , but got %d", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card of Ace of spades , but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected last card of Clubs  , but got %v", d[len(d)-1])

	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFIle("_decktesting")

	if len(loadedDeck) != 16 {
		// t.Error("Expected deck lenght of 16 , but got", len(d))
		t.Errorf("Expected deck lenght of 16 , but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
