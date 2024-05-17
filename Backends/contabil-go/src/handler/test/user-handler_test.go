package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	// Test your function
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user", nil)

}

func TestCreateUser(t *testing.T) {
	// Test your function
}

func TestUpdateUser(t *testing.T) {
	// Test your function
}

func TestDeleteUser(t *testing.T) {
	// Test your function
}
