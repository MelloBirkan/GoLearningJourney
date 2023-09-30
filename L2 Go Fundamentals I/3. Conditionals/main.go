package main

import "fmt"

func main() {
	var language = "Go"

	if language == "Java" {
		fmt.Println("Language is Java")
	} else if language == "C" {
		fmt.Println("Language is C")
	} else if language == "C++" {
		fmt.Println("Language is C++")
	} else {
		fmt.Println("Can't find the language")
	}
}