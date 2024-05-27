package handler

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

// GetUsers retrieves all users from both databases
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get users from database

	users, err := model.GetUsers()

	if err != nil {
		logging.FailedToFindOnDB("All Users", "", err.Inner, "handler.GetUsers")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to get users")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}

}

// CreateUser creates a user in both databases
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(domain.User)

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		logging.FailedToParseBody(err, "handler.CreateUser")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to parse body")
		return
	}

	result, tagErr := model.CreateUser(user)

	if tagErr != nil {
		logging.GenericError("Fail to create user", tagErr.Inner, "handler.CreateUser")
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}
}

// UpdateUser updates a user by id in both databases
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user from request
	user := new(domain.User)

	json.NewDecoder(r.Body).Decode(user)

	if user.ID == "" {
		logging.GenericError("Empty request body", nil, "handler.updateUser")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty request body")
		return
	}

	result, err := model.UpdateUser(user)

	if err != nil {
		logging.GenericError("Failed to update user", err.Inner, "handler.UpdateUser")
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}
}

// DeleteUser deletes a user by id in both databases
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get user from request
	id := r.PathValue("id")

	if id == "" {
		logging.GenericError("Empty id passed.", nil, "handler.DeleteUser")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty id passed.")
		return
	}

	err := model.DeleteUser(id)

	if err != nil {
		logging.GenericError("Failed to delete user", err.Inner, "handler.DeleteUser")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to delete user")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "User deleted successfully")
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Get user from request
	user := new(domain.User)
	json.NewDecoder(r.Body).Decode(user)

	if user.Name == "" || user.Password == "" {
		logging.GenericError("Empty request body", nil, "handler.Login")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty request body")
		return
	}

	result, err := model.LoginUser(*user)

	if err != nil {
		logging.GenericError("Failed to login user", err.Inner, "handler.handler.Login")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to login user")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
		return
	}

}
