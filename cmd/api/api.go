package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start() error {
	var router *chi.Mux = chi.NewRouter()

	var subrouter *chi.Mux = chi.NewRouter()

	router.Mount("/api/v1", subrouter) // create sub routing with prefix pat pattern

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.addr),
		Handler: router,
	}

	fmt.Printf("Server is listening on :%v .....\n\n", s.addr)

	fmt.Println("Press Ctrl+C to stop the server")

	return server.ListenAndServe()
}
