package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/logiXbomb/lenslocked.com/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.tmpl"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		log.Fatalf("-- could not render new users view: %v", err)
	}
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Fatalf("-- could not parse create user form: %v", err)
	}
	fmt.Fprint(w, form)
}
