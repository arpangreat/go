package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Swastik is a good guy")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Swastik is a student of class 12. He is a beast at coding. And he loves to play football")
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":9000", nil)
}
