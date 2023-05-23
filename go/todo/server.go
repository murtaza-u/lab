package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

type Server struct {
	router chi.Router
}

var session *Store

func init() {
	session = &Store{
		Mutex: sync.Mutex{},
		list:  make(map[int]*Todo),
	}
}

func NewServer() *Server {
	r := chi.NewRouter()
	r.Use(httprate.LimitByRealIP(20, time.Minute))

	r.Route("/todo", func(r chi.Router) {
		r.Post("/", Add)
		r.Get("/", GetAll)
		r.Get("/{id}", GetOne)
		r.Put("/{id}", ToggleCheck)
		r.Delete("/{id}", Delete)
	})

	return &Server{r}
}

func (s *Server) Run() error {
	return http.ListenAndServe(":5000", s.router)
}
