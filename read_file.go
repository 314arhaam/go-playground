package main

import (
	"fmt"
	"os"
)

func main() {
	text, err := os.ReadFile("tmp.txt")
	if (err == nil){
		fmt.Println(string(text))
		return
	} else {
		fmt.Println("error: ", err)
	}
}