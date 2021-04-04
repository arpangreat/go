package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	 <h1>hey , There</h1>
	`)
	fmt.Fprintf(w, "<p>You can %s add %s into it</p>", "even", "<strong>variables</strong>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8080", nil)

}
