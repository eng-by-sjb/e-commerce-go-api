package user

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Post("/register", h.handleRegister)
	router.Post("/login", h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json data from payload and unmarshal json data
	// validate json data
	// check if user already exists
	// create user

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
