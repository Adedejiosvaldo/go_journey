package main

import (
	"fmt"
	"fulltimego/modules/types"
	"fulltimego/modules/util"
)

func main() {
	user := types.User{Username: "James", Age: util.GetNumbers()}
	fmt.Println(user)
}
