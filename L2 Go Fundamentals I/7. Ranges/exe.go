package main

import "fmt"

func reduce(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	// Uncomment the line below after implementing your reduce() function
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
