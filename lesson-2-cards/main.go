package main

import "fmt"

// func main() {
// 	var varLong string = "Long Form of Variable Declaration"
// 	// Short form on var delcaration
// 	varShort := "Short form of variable declaration"
// 	// resassignin a variable
// 	varShort = "Reassigned var"

//		fmt.Println(varLong)
//		fmt.Println(varShort)
//	}

func main() {
	// Slice of strings
	cards := []string{"Ace of Diamonds", newCard()}

	// Append to a slice
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds "
}
