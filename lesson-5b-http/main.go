package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	bs := make([]byte, 999999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	// fmt.Println(resp)
}

func grew() {}
