package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	router := chi.NewRouter()

	sv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: sv}

	return &server, nil
}

//configs when the api is close
func (serv *Server) Close() error {
	// on here i will config the closures
	return nil
}

/*Method for start the server*/
func (serv *Server) Start() {
	log.Printf("Server on port:", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
