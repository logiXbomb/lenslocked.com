package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/logiXbomb/lenslocked.com/controllers"
)

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>What's going on?</h2>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "Oops... you didn't something wrong")
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	r.Handle("/", staticC.HomeView).Methods(http.MethodGet)
	r.Handle("/contact", staticC.ContactView).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.New).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.Create).Methods(http.MethodPost)
	r.HandleFunc("/faq", faq).Methods(http.MethodGet)

	r.NotFoundHandler = http.HandlerFunc(notFound)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("-- could not listen on port: %v", err)
	}
}

func must(err error) {
	if err != nil {
		log.Fatalf("-- cannot have an error: %v", err)
	}
}
