package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	/*
	   // Using Reader Interface
	   // Explanataion : Takes a info and sends it out
	*/
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// fmt.Println(resp)

	io.Copy(os.Stdout, resp.Body)
}

func grew() {}
