package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
