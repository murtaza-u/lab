package main

import (
	"log"
	"net/http"

	"github.com/murtaza-u/lab/go/todo/view"
	"github.com/murtaza-u/lab/go/todo/view/layout"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	c := layout.Base(view.Index())
	http.Handle("/", templ.Handler(c))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
