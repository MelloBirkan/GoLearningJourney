package main

import (
	"fmt"
	"net/http"
)

func showContactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./L4. Building for the Web/4. Serving HTML/static/index.html")
}

func main() {
	http.HandleFunc("/", showContactPage)
	// "FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root"
	// See: https://pkg.go.dev/net/http#FileServer

	// Tambem Funciona
	//fileServer := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fileServer)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
