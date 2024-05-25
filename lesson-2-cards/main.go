package main

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
	cards := newDeck()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds "
}
