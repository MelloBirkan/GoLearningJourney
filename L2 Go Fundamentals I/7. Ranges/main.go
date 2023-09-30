package main

import "fmt"

func main() {
	ahtletes := []string{"Ronaldo", "Messi", "Neymar"}
	nums := [4]int{1, 2, 4, 5}

	for i, ahtlete := range ahtletes {
		fmt.Printf("i = %d, ahtlete = %s\n", i, ahtlete)
	}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%d é um número primo\n", num)
		}
	}
}
