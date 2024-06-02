package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	fmt.Println("Hello")
}

func (eb englishBot) getGreeting() string {
	return "Hi There"
}

func (eb spanishBot) getGreeting() string {
	return "Hola"
}
