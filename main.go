package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Waffles are great!!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please get in touch <a href=\"mailto:support@test.com\">Support</a>.")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>What's going on?</h2>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "Oops... you didn't something wrong")
}

func main() {
	r := httprouter.New()

	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)

	r.NotFound = http.HandlerFunc(notFound)

	http.ListenAndServe(":3000", r)
}
