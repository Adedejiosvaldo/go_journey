package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// Embedding a struct
	alex := Person{firstName: "Alex", lastName: "OnGod", contactInfo: contactInfo{email: "Adedej", zip: 10000}}

	// Memeory reference to the alwx struct in memeory
	// Long approach
	// jimPointer := &alex
	// jimPointer.updateName("Jimmy")
	// alex.print()

	// Implementing a shortcut
	alex.updateName("Jimmy")
	alex.print()
	// okc := Person{firstName: alex.lastName, lastName: alex.lastName}

	// Upating a struct

}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

// Can be called by pointer with a type of perso
func (pointerToPerson *Person) updateName(firstName string) {
	// This is the operation to change the actual struct
	(*pointerToPerson).firstName = firstName
}
