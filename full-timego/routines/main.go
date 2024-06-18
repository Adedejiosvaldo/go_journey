package main

import (
	"fmt"
	"time"
)

// Channel would always block if it is full
func main() {
	resultch := make(chan string)

	go func() {
		result := <-resultch
		fmt.Println(result)

	}()

	resultch <- "food"
	// go fetch()
	// Syntax 1
	// go fetch(2)

	// syntax 2
	// go func() {
	// fetch()
	// }()
	// r := fetch()
	// fmt.Println(result)
}

func fetch(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprint("result %d", n)
}
