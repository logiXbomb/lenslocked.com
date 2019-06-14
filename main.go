package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/logiXbomb/lenslocked.com/controllers"
	"github.com/logiXbomb/lenslocked.com/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	must(contactView.Render(w, nil))
}

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
	homeView = views.NewView("bootstrap", "views/home.tmpl")
	contactView = views.NewView("bootstrap", "views/contact.tmpl")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods(http.MethodGet)
	r.HandleFunc("/contact", contact).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.New).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.Create).Methods(http.MethodPost)
	r.HandleFunc("/faq", faq).Methods(http.MethodGet)

	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		log.Fatalf("-- cannot have an error: %v", err)
	}
}
