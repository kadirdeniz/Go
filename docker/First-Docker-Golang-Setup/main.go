package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Started at 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Anasayfa")
	})

	http.ListenAndServe(":8080", nil)
}
