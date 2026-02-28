package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "EduScale API Running")
	})

	fmt.Println("Server started on :8080")

	http.ListenAndServe(":8080", nil)
}