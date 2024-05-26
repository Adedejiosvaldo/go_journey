package main

import "fmt"

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
