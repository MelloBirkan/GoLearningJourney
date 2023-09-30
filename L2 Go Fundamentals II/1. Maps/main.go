package main

import "fmt"

func main() {
	dictionary := map[string]string{
		"Go":     "Programming language.",
		"Gopher": "Software engineer who builds with Go.",
		"Golang": "Another name for Go.",
	}

	fmt.Println(dictionary)
	fmt.Println(dictionary["Gopher"])

	dictionary["Gopher"] = "The fuzzy mascot for Go."
	dictionary["Map"] = "An unordered data structure with key-value pairs"
	fmt.Println(dictionary)
}
