package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dev-by-sjb/e-commerce-go-api/types"
	"github.com/go-chi/chi"
)

func TestUserServiceHandler(t *testing.T) {
	// setup
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// act
	t.Run("Should fail if user payload is invalid", func(t *testing.T) {
		// expected
		expected := http.StatusBadRequest

		payload := types.RegisterUserPayload{
			FirstName: "Paul",
			LastName:  "Walker",
			Email:     "paul",
			Password:  "12345",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		// create a  http request recorder to simulate a request
		rr := httptest.NewRecorder()
		router := chi.NewRouter()

		router.Post("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		log.Println(rr.Code, rr.Body)

		// assert
		if expected != rr.Code {
			t.Errorf("expected %d, got %d", expected, rr.Code)
		}
	})
}

// mockUserStore implements the UserStore interface setup
type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u *types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int64) (*types.User, error) {
	return nil, nil
}
