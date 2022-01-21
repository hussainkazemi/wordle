package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Enter your own word: ")
	// var word string
	// fmt.Scanln(&word)
	// fmt.Printf("your word is %s", word)

	if isLoaded := LoadWords(); !isLoaded {
		checkError(fmt.Errorf("Cant load words "))
	}

}
