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
	var payload types.RegisterUserPayload
	var err error

	defer r.Body.Close()

	err = utils.ParseJSON[*types.RegisterUserPayload](r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload fields
	if err := utils.Validate.Struct(payload); err != nil {
		var errors string
		for _, err := range err.(validator.ValidationErrors) {
			errors += fmt.Sprintf("%s field failed validation check:- %s %s. ", err.Field(), err.Tag(), err.Param())
		}

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	// create user

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
