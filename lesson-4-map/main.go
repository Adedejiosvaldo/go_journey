package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "0xxhhrr",
	}

	printKeys(colors)

	//  Manipulating Maps
	// var colorsForMap
	colorsForMap := make(map[string]string)
	colorsForMap["red"] = "0x0000"

	// Deletes
	delete(colorsForMap, "red")

	// fmt.Println(colorsForMap)
}

func printKeys(m map[string]string) {

	for color, card := range m {
		fmt.Print("Key: " + color + ", value " + card)
	}
}
