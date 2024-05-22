package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
)

type TestUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type TestUserId struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func TestGetUsers(t *testing.T) {

	start()

	defer database.CloseConnections()

	// rr := httptest.NewRecorder()
	// req := httptest.NewRequest(http.MethodGet, "/user", nil)

}

func TestCreateUser(t *testing.T) {

	start()

	defer database.CloseConnections()

	body := TestUser{
		Name:     "Teste",
		Password: "123",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(body)

	if err != nil {
		util.LogHandler("Erro ao criar requisição", err, "TestCreateUser")
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/user", &b)

	CreateUser(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	response, err := io.ReadAll(rr.Body)

	if err != nil {
		util.LogHandler("Erro ao ler resposta", err, "TestCreateUser")
		t.Fatalf("Erro ao ler resposta")
	}

	expected := domain.UserDTO{
		Name: "Teste",
		Role: "ROLE_USER",
	}

	var result domain.UserDTO
	err = json.Unmarshal(response, &result)

	if err != nil {
		util.LogHandler("Erro ao deserializar resposta", err, "TestCreateUser")
		t.Fatalf("Erro ao deserializar resposta")
	}

	if result.Name != expected.Name {
		t.Errorf("handler returned wrong name: got %v want %v",
			result.Name, expected.Name)
	}

	if result.Role != expected.Role {
		t.Errorf("handler returned wrong role: got %v want %v",
			result.Role, expected.Role)
	}

	err = model.DeleteUser(result.ID)

	if err != nil {
		t.Errorf("Error while cleaning databases")
	}

}

func TestUpdateUser(t *testing.T) {
	start()

	defer database.CloseConnections()

	example := domain.User{
		Name:     "testUpdate",
		Password: "123",
	}

	userInBase, err := model.CreateUser(&example)

	if err != nil {
		util.LogHandler("Fail to create user for update test", err, "TestUpdateUser")
		t.Errorf("Fail to create user for update test")
	}

	defer model.DeleteUser(userInBase.ID)

	body := TestUserId{
		Id:   userInBase.ID,
		Name: "newUserName",
	}

	var b bytes.Buffer

	err = json.NewEncoder(&b).Encode(body)

	if err != nil {
		util.LogHandler("Erro ao criar requisição", err, "TestCreateUser")
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/user", &b)

	UpdateUser(rr, req)

}

func TestDeleteUser(t *testing.T) {
	// Test your function
}

func start() {
	err := config.Load()

	if err != nil {
		util.LogHandler("Erro ao carregar as configurações", err, "TestCreateUser")
	}

	err = database.OpenConnectionLocal()

	if err != nil {
		util.LogHandler("Erro ao conectar ao banco de dados local", err, "main")
	}

	err = database.OpenConnectionCloud()

	if err != nil {
		util.LogHandler("Erro ao conectar ao banco de dados em nuvem", err, "main")
	}
}
