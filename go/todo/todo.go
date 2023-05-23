package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos := session.GetAll()
	out, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	_id := chi.URLParam(r, "id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		http.Error(w, fmt.Errorf(
			"failed to parse id, only int is allowed: %w", err).Error(),
			http.StatusBadRequest)
		return
	}

	if !session.Exists(id) {
		http.Error(w, "id does not exists", http.StatusBadRequest)
		return
	}

	todo := session.Get(id)
	out, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func Add(w http.ResponseWriter, r *http.Request) {
	req := new(Todo)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "missing text", http.StatusBadRequest)
		return
	}

	session.Put(req)
	out, err := json.Marshal(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func ToggleCheck(w http.ResponseWriter, r *http.Request) {
	_id := chi.URLParam(r, "id")
	id, err := strconv.Atoi(_id)

	if !session.Exists(id) {
		http.Error(w, "id does not exists", http.StatusBadRequest)
		return
	}

	t := session.Get(id)
	t.Checked = !t.Checked

	session.Put(t)
	out, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	_id := chi.URLParam(r, "id")
	id, err := strconv.Atoi(_id)

	if !session.Exists(id) {
		http.Error(w, "id does not exists", http.StatusBadRequest)
		return
	}

	t := session.Get(id)
	out, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Del(id)
	w.Write(out)
}
