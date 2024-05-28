package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) shuffle() {
	// Uniuq time seeed
	time := time.Now().UnixNano()
	fmt.Println(time)

	// Creating a new source
	source := rand.NewSource(time)

	// Source is used as a basis for our random generator
	r := rand.New(source)
	// Not
	for index := range d {
		// Len -> lenght of a slice
		newPosition := r.Intn(len(d) - 1)

		// Take whatever is at newPosition and assign it to d[index]
		// Take whatever is at d[index], and assign it to d[newPosition]
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
