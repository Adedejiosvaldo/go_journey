package main

// func main() {
// 	var varLong string = "Long Form of Variable Declaration"
// 	// Short form on var delcaration
// 	varShort := "Short form of variable declaration"
// 	// resassignin a variable
// 	varShort = "Reassigned var"

//		fmt.Println(varLong)
//		fmt.Println(varShort)
// fmt.Println(cards.toString())
//	}

func main() {
	// Slice of strings
	// Generates a new deck of cards
	cards := newDeck()
	// Deal cards
	// dealtCards := deal(cards, 4)

	// Save to file
	// cards.saveToFile("my_deck")

	cards.shuffle()
	// loadedCards := newDeckFromFIle("my_deck")
	cards.print()

}
