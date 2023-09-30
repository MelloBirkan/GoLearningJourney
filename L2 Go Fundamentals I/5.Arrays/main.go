package main

import "fmt"
func main() {
	// fibonacci 0 1 1 2 3 5 8 13
	// one way to declare an array
	var fibonacci [8]int
	fibonacci[0] = 0
	fibonacci[1] = 1
	fibonacci[2] = 1
	fibonacci[3] = 2
	fibonacci[4] = 3
	fibonacci[5] = 5
	fibonacci[6] = 8
	fibonacci[7] = 13
	fmt.Println(fibonacci)

	// another way to declare an array
	fibonacci2 := [8]int{0, 1, 1, 2, 3, 5, 8, 13}
	fmt.Println(fibonacci2)

	// Size of an array
	fmt.Println("This array have:", len(fibonacci))

	// accessing an array
	fmt.Println(fibonacci[0])
	fmt.Println(fibonacci2[7])
	fmt.Println(fibonacci[2:5])
	
}