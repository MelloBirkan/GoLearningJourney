package main

import "fmt"

func main() {
	var languages = []string{"JavaScript", "Ruby", "Python"}
	fmt.Println(languages)
	fmt.Println(len(languages))
	fmt.Println(languages[0])
	fmt.Println(languages[1:3])
	languages = append(languages, "PHP")
	fmt.Println(languages)
}
