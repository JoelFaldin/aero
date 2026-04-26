package main

import (
	"fmt"
	"log"
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the server!")
}

func main() {
	http.HandleFunc("/", basicHandler)

	fmt.Println("Server starting at http://localhost:8080!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error:", err)
	}
}
