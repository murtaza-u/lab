package server

import "net/http"

const phrase = "In loving memories of Meghan Vinayak Mane"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(phrase))
}

func New() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":8080", nil)
}
