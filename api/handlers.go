package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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

func (s *Server) getAllBlocksHandler(w http.ResponseWriter, r *http.Request) {
	// Request all blocks from NATS
	msg, err := s.nc.Request("block.registry.getall", nil, 2*time.Second)
	if err != nil {
		log.Printf("Error requesting blocks: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check if message data is not empty
	if len(msg.Data) == 0 {
		http.Error(w, "No blocks found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(msg.Data)
}

func StartHTTPServer(nc *nats.Conn) {
	server := NewServer(nc)
	server.AddRoute("/api/blocks/new", server.createBlockHandler, http.MethodPost)
	server.AddRoute("/api/blocks/all", server.getAllBlocksHandler, http.MethodGet)
	server.Start()
}
