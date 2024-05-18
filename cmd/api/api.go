package api

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/dev-by-sjb/e-commerce-go-api/service/user"
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

	var v1Router *chi.Mux = chi.NewRouter()

	router.Mount("/api/v1", v1Router) // create sub routing with prefix to handle api versioning

	// Add user services that takes in the subrouter
	var userStore = user.NewStore(s.db)
	var userHandler = user.NewHandler(userStore)
	userHandler.RegisterRoutes(v1Router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.addr),
		Handler: router,
	}

	fmt.Printf("Server is listening on :%v .....\n\n", s.addr)

	fmt.Print("Press Ctrl+C to stop the server\n\n")

	return server.ListenAndServe()
}
