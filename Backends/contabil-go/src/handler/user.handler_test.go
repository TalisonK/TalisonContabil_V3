package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
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

func TestMain(m *testing.M) {

	err := config.Load()

	if err != nil {
		logging.GenericError("Erro ao carregar configurações", err, "handler.Testmain")
		return
	}

	err = database.OpenConnectionLocal()

	if err != nil {
		logging.FailedToOpenConnection("local", err, "handler.Testmain")
		return
	}

	err = database.OpenConnectionCloud()

	if err != nil {
		logging.FailedToOpenConnection("cloud", err, "handler.Testmain")
		return
	}

	exitVal := m.Run()

	defer database.CloseConnections()

	os.Exit(exitVal)
}

func TestGetUsers(t *testing.T) {

	// rr := httptest.NewRecorder()
	// req := httptest.NewRequest(http.MethodGet, "/user", nil)

}

func TestCreateUser(t *testing.T) {

	body := TestUser{
		Name:     "Teste",
		Password: "123",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(body)

	if err != nil {
		logging.GenericError("Erro ao criar requisição", err, "TestCreateUser")
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
		logging.GenericError("Erro ao ler resposta", err, "TestCreateUser")
		t.Fatalf("Erro ao ler resposta")
	}

	expected := domain.UserDTO{
		Name: "Teste",
		Role: "ROLE_USER",
	}

	var result domain.UserDTO
	err = json.Unmarshal(response, &result)

	if err != nil {
		logging.GenericError("Erro ao deserializar resposta", err, "TestCreateUser")
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

	tagErr := model.DeleteUser(result.ID)

	if tagErr != nil {
		t.Errorf("Error while cleaning databases")
	}

}

func TestUpdateUser(t *testing.T) {

	example := domain.User{
		Name:     "testUpdate",
		Password: "123",
	}

	userInBase, tagErr := model.CreateUser(&example)

	if tagErr != nil {
		logging.GenericError("Fail to create user for update test", tagErr.Inner, "TestUpdateUser")
		t.Errorf("Fail to create user for update test")
	}

	defer model.DeleteUser(userInBase.ID)

	body := TestUserId{
		Id:   userInBase.ID,
		Name: "newUserName",
	}

	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(body)

	if err != nil {
		logging.GenericError("Erro ao criar requisição", err, "TestUpdateUser")
		t.Fatalf("Erro ao criar requisição")
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/user", &b)

	UpdateUser(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	response, err := io.ReadAll(rr.Body)

	if err != nil {
		logging.GenericError("Erro ao ler resposta", err, "TestCreateUser")
		t.Fatalf("Erro ao ler resposta")
	}

	expected := domain.UserDTO{
		Name: "newUserName",
	}

	var result domain.UserDTO
	err = json.Unmarshal(response, &result)

	if err != nil {
		logging.GenericError("Erro ao deserializar resposta", err, "TestCreateUser")
		t.Fatalf("Erro ao deserializar resposta")
	}

	if result.Name != expected.Name {
		t.Errorf("handler returned wrong name: got %v want %v",
			result.Name, expected.Name)
	}

}

func TestDeleteUser(t *testing.T) {

	//TOFIX
	// example := domain.User{
	// 	Name:     "testUpdate",
	// 	Password: "123",
	// }

	// userInBase, err := model.CreateUser(&example)

	// if err != nil {
	// 	util.LogHandler("Fail to create user for update test", err, "TestUpdateUser")
	// 	t.Errorf("Fail to create user for update test")
	// }

	// u, err := url.Parse("/user")

	// if err != nil {
	// 	util.LogHandler("Fail to parse url", err, "TestDeleteUser")
	// 	t.Errorf("Fail to parse url")
	// }

	// q := u.Query()
	// q.Set("id", userInBase.ID)
	// u.RawQuery = q.Encode()

	// rr := httptest.NewRecorder()
	// req := httptest.NewRequest(http.MethodDelete, u.String(), nil)

	// DeleteUser(rr, req)

	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

}
