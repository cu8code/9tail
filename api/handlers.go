package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cu8code/9tail/antlr/parser"
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
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the file has .ninetail extension
	if !strings.HasSuffix(header.Filename, ".ninetail") {
		http.Error(w, "File must have .ninetail extension", http.StatusBadRequest)
		return
	}

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	str := string(content)
	parser.Parse(antlr.NewInputStream(str))

	// Set the response headers
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", header.Filename))

	// Write the file content to the response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(content)

	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
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
