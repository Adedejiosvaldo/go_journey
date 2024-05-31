package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	alex := Person{"Alex", "Anderson"}

	okc := Person{firstName: alex.lastName, lastName: alex.lastName}
	fmt.Println(okc)

}
