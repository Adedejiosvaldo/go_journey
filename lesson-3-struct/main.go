package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	alex := Person{"Alex", "Anderson"}

	fmt.Println(alex)
	okc := Person{firstName: alex.lastName, lastName: alex.lastName}

	// Upating a struct
	okc.firstName = "The Name Changed"

	fmt.Println(okc)

}
