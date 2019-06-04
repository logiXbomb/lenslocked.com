package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("./exp/hello.tmpl")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name   string
		Great  string
		Things []string
		Count  int
	}{
		Name:  "waffles, and stuff",
		Great: "waffles are",
		Count: 25,
		Things: []string{
			"banana",
			"syrup",
		},
	}

	err = t.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
