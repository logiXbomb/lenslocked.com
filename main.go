package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please get in touch <a href=\"mailto:support@test.com\">Support</a>.")
	} else {
		fmt.Fprint(w, "<h1>Waffles are great!!</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)

	http.ListenAndServe(":3000", nil)
}
