package main

import "fmt"

// Create new type of deck
// which is a slice of strings

type deck []string

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
