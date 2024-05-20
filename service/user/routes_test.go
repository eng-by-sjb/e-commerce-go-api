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

type testCase struct {
	name     string
	path     string
	method   string
	payload  types.RegisterUserPayload
	expected int
}

var testCases = []testCase{
	{
		name:   "should fail to register new user if payload is invalid",
		path:   "/register",
		method: http.MethodPost,
		payload: types.RegisterUserPayload{
			FirstName: "Paul",
			LastName:  "Walker",
			Email:     "paul.com",
			Password:  "123",
		},
		expected: http.StatusBadRequest,
	},
	{
		name:   "should register new user successfully",
		path:   "/register",
		method: http.MethodPost,
		payload: types.RegisterUserPayload{
			FirstName: "Paul",
			LastName:  "Walker",
			Email:     "paulwalker@gmail.com",
			Password:  "123456",
		},
		expected: http.StatusCreated,
	},
}

func TestUserServiceHandler(t *testing.T) {
	// setup
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload, _ := json.Marshal(tc.payload)

			req, err := http.NewRequest(http.MethodPost, tc.path, bytes.NewBuffer(payload))
			if err != nil {
				t.Fatal(err)
			}

			// create a new request recorder to simulate a request
			rr := httptest.NewRecorder()
			router := chi.NewRouter()

			router.MethodFunc(tc.method, tc.path, handler.handleRegister)
			router.ServeHTTP(rr, req)

			log.Println(rr.Code, rr.Body)

			// assert
			if tc.expected != rr.Code {
				t.Errorf("expected %d but got %d", tc.expected, rr.Code)
			}
		})
	}
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
