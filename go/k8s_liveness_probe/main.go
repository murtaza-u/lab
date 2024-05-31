package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type I struct {
	mutex sync.Mutex
	count uint
}

var i I

func handler(w http.ResponseWriter, r *http.Request) {
	if i.count >= 5 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	i.mutex.Lock()

	i.count++
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Count: %d\n", i.count)

	i.mutex.Unlock()
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on port :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
