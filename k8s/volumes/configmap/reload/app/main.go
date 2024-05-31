package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

var Foo string

const EnvDir = "env/foo"

func readEnv() {
	godotenv.Overload("env/foo")
	Foo = os.Getenv("FOO")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo: %s\n", Foo)
}

func watch(w *fsnotify.Watcher) {
	for {
		select {
		case e, ok := <-w.Events:
			if !ok {
				log.Println("!ok. Terminating watcher")
				return
			}

			log.Printf("An event occured: %s\n", e.String())

			if e.Op == fsnotify.Remove {
				if err := w.Add(EnvDir); err != nil {
					log.Printf("error: %s\n", err.Error())
					return
				}

				readEnv()
				continue
			}

			if e.Op&fsnotify.Write == fsnotify.Write {
				readEnv()
			}

		case err, ok := <-w.Errors:
			if !ok {
				log.Println("!ok. Terminating watcher")
				return
			}

			log.Printf("fsnotify[err]: %s\n", err.Error())
			return
		}
	}
}

func main() {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	err = w.Add(EnvDir)
	if err != nil {
		log.Fatal(err)
	}

	go watch(w)
	readEnv()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
