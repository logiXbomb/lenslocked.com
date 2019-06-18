package controllers

import "github.com/logiXbomb/lenslocked.com/views"

func NewStatic() *Static {
	return &Static{
		HomeView:    views.NewView("bootstrap", "views/static/home.tmpl"),
		ContactView: views.NewView("bootstrap", "views/static/contact.tmpl"),
	}
}

type Static struct {
	HomeView    *views.View
	ContactView *views.View
}
