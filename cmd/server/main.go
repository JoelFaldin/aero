package main

import (
	"fmt"
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the server!")
}

func healthChecker(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc("/", basicHandler)
	http.HandleFunc("/health", healthChecker)

	startPort := 8080
	for port := startPort; port < startPort+10; port++ {
		addr := fmt.Sprintf(":%d", port)

		fmt.Println("Server starting on port", port)

		err := http.ListenAndServe(addr, nil)
		if err == nil {
			fmt.Printf("Server started off port %d!\n", port)
			return
		}

		fmt.Printf("Port %d is occupied, trying next...\n", port)
	}
}
