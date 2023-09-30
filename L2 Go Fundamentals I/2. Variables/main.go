package main
import "fmt"

func main() {
	// Forma completa
	var product string = "T-shirt"
	// Forma curta
	var cost = 20
	// Forma mais curta
	quantity := 5
	// Constantes
	const pi = 3.14159265359

	fmt.Println("Product's value is:", product)
	fmt.Printf("Cost's type is: %T\n", cost)

	quantity = 19
	fmt.Println("The quantity of products is:", quantity)

	fmt.Println("The value of pi is:", pi)
}