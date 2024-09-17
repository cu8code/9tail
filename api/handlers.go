package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

type Server struct {
	mux *http.ServeMux
	nc  *nats.Conn
}

type BlockRequest struct {
	Type string `json:"type"`
}

func NewServer(nc *nats.Conn) *Server {
	return &Server{
		mux: http.NewServeMux(),
		nc:  nc,
	}
}

func (s *Server) AddRoute(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if r.Method != method {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}
		}
		handler(w, r)
	})
}

func (s *Server) Start() {
	log.Println("Starting HTTP server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.mux))
}

func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NATS-based Workflow System - HTTP Server"))
}

func getBlocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blocks endpoint"))
}

func (s *Server) createBlockHandler(w http.ResponseWriter, r *http.Request) {
	var blockReq BlockRequest
	err := json.NewDecoder(r.Body).Decode(&blockReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Creating block of type %s\n", blockReq.Type)

	// Publish block registration request to NATS
	err = s.nc.Publish("block.registry", []byte(blockReq.Type))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Block created"))
}

func StartHTTPServer(nc *nats.Conn) {
	server := NewServer(nc)
	server.AddRoute("/api", randomNumberHandler, http.MethodGet)
	server.AddRoute("/api/blocks", getBlocksHandler, http.MethodGet)
	server.AddRoute("/api/blocks/new", server.createBlockHandler, http.MethodPost)
	server.Start()
}
