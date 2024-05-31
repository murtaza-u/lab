package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

const svcname = "datastore.default.svc.cluster.local"

var dataF, port, hostname string

func init() {
	dataF = os.Getenv("DATAFILE")
	if dataF == "" {
		dataF = "node.data"
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	hostname, _ = os.Hostname()
}

func main() {
	r := chi.NewRouter()
	r.Get("/", handleGet)
	r.Get("/peers", handlePeerGet)
	r.Put("/", handlePut)
	http.ListenAndServe(":"+port, r)
}

func handlePeerGet(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(dataF)
	if err != nil {
		log.Printf("os.ReadFile: %s\n", err.Error())
		return
	}

	w.Write(data)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer r.Body.Close()
	if body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, err := os.OpenFile(dataF, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data = append(data, []byte("\n")...)
	_, err = f.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("data stored in data store\n"))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	peers, err := getPeers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data []byte
	for _, p := range peers {
		resp, err := http.Get("http://" + p + "/peers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data = append(data, body...)
	}

	w.Write(data)
}

func getPeers() ([]string, error) {
	cname, addrs, err := net.LookupSRV("", "", svcname)
	if err != nil {
		log.Printf("ERR: %s\n", err.Error())
		return nil, err
	}
	log.Printf("CNAME: %s\n", cname)

	var peers []string
	for _, a := range addrs {
		log.Printf(
			"target: %s | port: %d | weight: %d | priority: %d\n",
			a.Target, a.Port, a.Weight, a.Priority,
		)
		peers = append(peers, fmt.Sprintf("%s:%d", a.Target, a.Port))
	}

	return peers, nil
}
