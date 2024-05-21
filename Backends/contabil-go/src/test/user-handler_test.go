package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/go-chi/chi/v5"
)

func TestGetUsers(t *testing.T) {

	r := chi.NewRouter()

	r.Get("/user", handler.GetUsers)

	req, err := http.NewRequest(http.MethodGet, "/user", nil)

	if err != nil {
		util.LogHandler("Erro ao criar requisição", err, "TestGetUsers")
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateUser(t *testing.T) {
	r := chi.NewRouter()

	r.Get("/user", handler.GetUsers)

	user := []byte(`{"name": "teste", "password":"123456"}`)

	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(user))

	if err != nil {
		util.LogHandler("Erro ao criar requisição", err, "TestCreateUser")
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	var body *domain.UserDTO

	json.NewDecoder(rr.Body).Decode(body)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateUser(t *testing.T) {
	// Test your function
}

func TestDeleteUser(t *testing.T) {
	// Test your function
}
