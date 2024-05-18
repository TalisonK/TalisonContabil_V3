package handler

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/gofiber/fiber/v3"
	"github.com/magiconair/properties/assert"
)

func TestGetUsers(t *testing.T) {

	testApp := fiber.New()

	testApp.Get("/user", handler.GetUsers)

	req, err := http.NewRequest(http.MethodGet, "/user", nil)

	if err != nil {
		util.LogHandler("Erro ao criar requisição", err, "TestGetUsers")
		t.Fatalf("Erro ao criar requisição")
	}

	resp, err := testApp.Test(req)

	if err != nil {
		util.LogHandler("Erro ao testar rota", err, "TestGetUsers")
		t.Fatalf("Erro ao testar rota")
	}

	assert.Equal(t, resp.StatusCode, 200)

}

func TestCreateUser(t *testing.T) {
	// Test your function

	testApp := fiber.New()

	testApp.Post("/user", handler.CreateUser)

	user := []byte(`{"name": "Talison", "password":"123456"}`)

	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(user))

	if err != nil {
		util.LogHandler("Erro ao criar requisição", err, "TestCreateUser")
		t.Fatalf("Erro ao criar requisição")
	}

	resp, err := testApp.Test(req)

	if err != nil {
		util.LogHandler("Erro ao testar rota", err, "TestCreateUser")
		t.Fatalf("Erro ao testar rota")
	}

	assert.Equal(t, resp.StatusCode, 200)
}

func TestUpdateUser(t *testing.T) {
	// Test your function
}

func TestDeleteUser(t *testing.T) {
	// Test your function
}
