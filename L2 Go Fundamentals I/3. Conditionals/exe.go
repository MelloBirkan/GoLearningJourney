package main

import "fmt"

func main() {
	// TODO
	var number int = 21

	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 100 && number > 0 {
		fmt.Println(number, "is positive")
	} else if number >= 100 {
		fmt.Println(number, "is positive and is a large number")
	}
}
