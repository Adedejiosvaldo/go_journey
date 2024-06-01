package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red": "0xxhhrr",
	// }
	// fmt.Println(colors)

	//  Manipulating Maps
	// var colorsForMap
	colorsForMap := make(map[string]string)
	colorsForMap["red"] = "0x0000"

	fmt.Println(colorsForMap)
}
