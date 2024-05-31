package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FOO"))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func serve(req *Request) Response {
	router := chi.NewRouter()
	router.Get("/", indexHandler)
	router.Get("/foo", fooHandler)
	router.Post("/", postHandler)
	w := new(ResponseWriter)
	r, err := http.NewRequest(req.Method, req.URL, bytes.NewReader(req.Body))
	if err != nil {
		return Response{
			Data:   []byte(err.Error()),
			Status: http.StatusBadRequest,
		}
	}
	router.ServeHTTP(w, r)
	return Response{
		Data:   w.buffer.Bytes(),
		Status: w.statusCode,
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := serve(&Request{
			Method: r.Method,
			URL:    r.URL.Path,
			Body:   body,
			Header: Header(r.Header),
		})
		w.WriteHeader(resp.Status)
		w.Write(resp.Data)
	})
	log.Printf("listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

type ResponseWriter struct {
	buffer     bytes.Buffer
	statusCode int
}

func (w *ResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *ResponseWriter) Write(b []byte) (n int, err error) {
	return w.buffer.Write(b)
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.statusCode = status
}
