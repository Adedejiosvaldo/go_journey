package main

import (
	"fmt"
	"os"
	"strings"
)

// Create new type of deck
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three", "Four"}

	for _, card := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card+" of "+value)
		}
	}

	return cards
}

func deal(cards deck, cardRange uint) (deck, deck) {
	// dealtCard := cards[0:cardRange]
	// notDealthWith := cards[cardRange:]
	return cards[:cardRange], cards[cardRange:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFIle(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// Bytes to string
	// string(bs) -> converts bytes to strings
	arrayOfString := strings.Split(string(bytes), ",")
	// converts []strings to type of dec
	return deck(arrayOfString)
}
