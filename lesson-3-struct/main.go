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

	jimPointer := &alex
	jimPointer.updateName("Jimmy")
	alex.print()

	// okc := Person{firstName: alex.lastName, lastName: alex.lastName}

	// Upating a struct

}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *Person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
