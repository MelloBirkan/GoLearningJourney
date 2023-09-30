package main

import "fmt"

func getARectangleArea(width, length int) string {

	var product = width * length

	if product < 50 {
		return fmt.Sprintf("The area is %d, which is less than 50", product)
	} else {
		return fmt.Sprintf("The area is %d, which is greater or equal to 50", product)
	}
}

func main() {
	// TODO
	fmt.Println(getARectangleArea(5, 10))
}
