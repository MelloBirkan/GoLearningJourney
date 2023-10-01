package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world  :]")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact Info</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	fmt.Println("Server started at:", time.Now())
	http.ListenAndServe(":3000", nil)
}
