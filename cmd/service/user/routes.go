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
	router.HandleFunc("/register", h.handleRegister)
	router.HandleFunc("/login", h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
