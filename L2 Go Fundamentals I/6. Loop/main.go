package main

import "fmt"

func main() {
	// Loops 1
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	// outra forma
	i := 0

	for i <= 20 {
		fmt.Println(i)
		i++
	}
}
