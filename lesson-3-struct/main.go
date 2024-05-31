package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := Person{firstName: "Alex", lastName: "OnGod", contact: contactInfo{email: "Adedej", zip: 10000}}

	fmt.Printf("%+v", alex)
	// okc := Person{firstName: alex.lastName, lastName: alex.lastName}

	// Upating a struct

}
