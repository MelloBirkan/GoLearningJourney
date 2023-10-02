package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=/cities>Hello stranger<a>")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<a href=/>Home</br><a>")
	fmt.Fprintln(w, "List of most populous cities:")
	for _, city := range cities {
		fmt.Fprintln(w, city)
	}
}

func main() {
	http.HandleFunc("/", index2)
	http.HandleFunc("/cities", cityList)
	fmt.Println("Server started!")
	http.ListenAndServe(":3000", nil)
}
