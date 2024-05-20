package user

import (
	"fmt"
	"net/http"

	"github.com/dev-by-sjb/e-commerce-go-api/types"
	"github.com/dev-by-sjb/e-commerce-go-api/utils"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
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

	// check if user already exists and check if email is already in use
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("email already in use"))
		return
	}

	hashedPassword, err := utils.CreateHashedPassword(payload.Password)
	if err != nil {
		// todo write error to app logs
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("something went wrong. try again later"))
	}

	// create user
	h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	utils.WriteJSON(w, http.StatusCreated, user)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
