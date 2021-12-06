package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server Started at 8091")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Anasayfa")
	})

	http.ListenAndServe(":8090", nil)

}
