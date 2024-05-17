package handler

import (
	"context"
	"fmt"
	"time"

	"encoding/json"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/model"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUsers retrieves all users from both databases
func GetUsers(c fiber.Ctx) error {
	// Get users from database
	if database.CheckCloudDB() {
		// Get users from cloud
		result, err := getCloudUsers()
		if err != nil {
			util.LogHandler("Failed to get users from cloud database.", err, "getUsers")
		} else {
			util.LogHandler(fmt.Sprintf("Users successfully retrieved %d rows from cloud database.", len(result)), nil, "getUsers")
			return c.Status(200).JSON(result)
		}
	}

	// Get users from local
	if database.CheckLocalDB() {
		result, err := getLocalUsers()
		if err != nil {
			util.LogHandler("Failed to get users from local database.", err, "getUsers")
		} else {
			util.LogHandler(fmt.Sprintf("Users successfully retrieved %d rows from local database.", len(result)), nil, "getUsers")
			return c.Status(200).JSON(result)
		}
	}

	util.LogHandler("No database connection available.", nil, "getUsers")
	return c.Status(500).JSON("No database connection available.")
}

// CreateUser creates a user in both databases
func CreateUser(c fiber.Ctx) error {
	// Get user from request
	user := new(model.User)
	raw := c.Body()

	err := json.Unmarshal(raw, user)

	if err != nil {
		util.LogHandler("Failed to parse body.", err, "createUser")
		return c.Status(400).JSON("Failed to parse body.")
	}

	user.Role = "ROLE_USER"

	var newId string = ""

	// Create user in cloud
	if database.CheckCloudDB() {

		userParse := bson.M{
			"name":      user.Name,
			"password":  user.Password,
			"role":      user.Role,
			"createdAt": primitive.NewDateTimeFromTime(time.Now()),
			"updatedAt": primitive.NewDateTimeFromTime(time.Now()),
		}

		result, err := database.DBCloud.User.InsertOne(context.Background(), userParse)
		if err != nil {
			util.LogHandler("Failed to create user in cloud database.", err, "createUser")
		} else {
			newId = result.InsertedID.(primitive.ObjectID).Hex()
			util.LogHandler(fmt.Sprintf("User %s successfully created in cloud database.", newId), nil, "createUser")
		}
	}

	if newId == "" {
		users, err := getLocalUsers()
		if err != nil {
			util.LogHandler("Failed to get users from local database.", err, "createUser")
			return c.Status(500).JSON("Failed to get users from local database.")
		}
		for {
			newId = primitive.NewObjectID().String()
			var exists bool = false
			for _, u := range users {
				if u.ID == newId {
					exists = true
					break
				}
			}
			if !exists {
				break
			}
		}
	}

	// Create user in local
	if database.CheckLocalDB() {
		user.ID = newId
		now := time.Now()
		user.CreatedAt = now.Format(time.RFC3339)
		user.UpdatedAt = now.Format(time.RFC3339)
		result := database.DBlocal.Create(user)
		if result.Error != nil {
			util.LogHandler("Failed to create user in local database.", result.Error, "createUser")
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully created in local database.", user.ID), nil, "createUser")
			return c.Status(201).JSON(user)
		}

	}

	if !database.CheckCloudDB() && !database.CheckLocalDB() {

		util.LogHandler("No database connection available.", nil, "createUser")
		return c.Status(500).JSON("No database connection available.")
	}
	return c.Status(500).JSON("Failed to create user.")
}

// UpdateUser updates a user by id in both databases
func UpdateUser(c fiber.Ctx) error {
	// Get user from request
	user := new(model.User)
	raw := c.Body()

	err := json.Unmarshal(raw, user)

	baseUser, err := findUserById(user.ID)

	if user.Name != "" {
		baseUser.Name = user.Name
	}
	if user.Password != "" {
		baseUser.Password = user.Password
	}
	if user.Role != "" {
		baseUser.Role = user.Role
	}
	baseUser.UpdatedAt = time.Now().Format(time.RFC3339)

	user = &baseUser

	if err != nil {
		util.LogHandler("Failed to find user.", err, "updateUser")
		return c.Status(400).JSON("Failed to find user.")
	}

	// Update user in cloud
	if database.CheckCloudDB() {

		userParse := userToPrim(*user)

		id, _ := primitive.ObjectIDFromHex(user.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": userParse}

		_, err = database.DBCloud.User.UpdateOne(context.Background(), filter, update)
		if err != nil {
			util.LogHandler("Failed to update user in cloud database.", err, "updateUser")
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully updated in cloud database.", user.ID), nil, "updateUser")
		}
	}

	// Update user in local
	if database.CheckLocalDB() {

		aux, _ := time.Parse(time.RFC3339, user.CreatedAt)
		user.CreatedAt = aux.Format(time.RFC3339)

		result := database.DBlocal.Save(user)
		if result.Error != nil {
			util.LogHandler("Failed to update user in local database.", result.Error, "updateUser")
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully updated in local database.", user.ID), nil, "updateUser")
			return c.Status(200).JSON(user)
		}
	}

	if !database.CheckCloudDB() && !database.CheckLocalDB() {
		util.LogHandler("No database connection available.", nil, "updateUser")
		return c.Status(500).JSON("No database connection available.")
	}

	return c.Status(500).JSON("Failed to update user.")
}

// DeleteUser deletes a user by id in both databases
func DeleteUser(c fiber.Ctx) error {
	// Get user from request
	id := c.Params("id")

	if id == "" {
		util.LogHandler("Empty id passed.", nil, "deleteUser")
		return c.Status(400).JSON("Failed to parse user ID.")
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		util.LogHandler(fmt.Sprintf("Failed to parse user ID %s.", id), err, "deleteUser")
		return c.Status(400).JSON("Failed to parse user ID.")
	}

	// Delete user in cloud
	if database.CheckCloudDB() {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.User.DeleteOne(context.Background(), filter)
		if err != nil {
			util.LogHandler("Failed to delete user in cloud database.", err, "deleteUser")
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully deleted in cloud database.", id), nil, "deleteUser")
		}
	}

	// Delete user in local
	if database.CheckLocalDB() {
		result := database.DBlocal.Delete(&model.User{}, "id = ?", id)
		if result.Error != nil {
			util.LogHandler("Failed to delete user in local database.", result.Error, "deleteUser")
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully deleted in local database.", id), nil, "deleteUser")
			return c.Status(200).JSON("User successfully deleted.")
		}
	}

	if !database.CheckCloudDB() && !database.CheckLocalDB() {
		util.LogHandler("No database connection available.", nil, "deleteUser")
		return c.Status(500).JSON("No database connection available.")
	}

	return c.Status(500).JSON("Failed to delete user.")
}

// findUserById retrieves a user by its ID
func findUserById(id string) (model.User, error) {
	var user model.User
	if database.CheckCloudDB() {
		idParse, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			util.LogHandler(fmt.Sprintf("Failed to parse user ID %s.", id), err, "findUserById")
			return user, err
		}

		filter := bson.M{"_id": idParse}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			util.LogHandler(fmt.Sprintf("User %s not found in cloud database.", id), nil, "findUserById")
		} else {
			user = primToUser(result)
			util.LogHandler(fmt.Sprintf("User %s successfully found in cloud database.", id), nil, "findUserById")
			return user, nil
		}
	}

	if database.CheckLocalDB() {
		result := database.DBlocal.First(&user, "id = ?", id)

		if result.Error != nil {
			util.LogHandler("Failed to find user in local database.", result.Error, "findUserById")
			return user, result.Error
			//TODO: equalize bases de users
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully found in local database.", id), nil, "findUserById")
			return user, nil
		}
	}

	util.LogHandler("No database connection available.", nil, "findUserById")
	return user, nil
}

// getCloudUsers retrieves all users from the cloud database
func getCloudUsers() ([]model.User, error) {
	colect, err := database.DBCloud.User.Find(context.TODO(), &bson.D{})

	result := []model.User{}

	//colect.All(context.Background(), &result)

	for colect.Next(context.Background()) {
		var user bson.M
		colect.Decode(&user)

		result = append(result, primToUser(user))
	}

	if err != nil {
		util.LogHandler("Failed to get users from cloud database.", err, "getUsers")
		return nil, err
	}
	return result, nil
}

// primToUser converts a primitive.M to a model.User
func primToUser(user primitive.M) model.User {
	var usuario model.User

	usuario.ID = user["_id"].(primitive.ObjectID).Hex()
	usuario.Name = user["name"].(string)
	usuario.Password = user["password"].(string)
	usuario.Role = user["role"].(string)
	usuario.CreatedAt = user["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	usuario.UpdatedAt = user["updatedAt"].(primitive.DateTime).Time().Format(time.RFC3339)

	return usuario
}

// userToPrim converts a model.User to a primitive.M
func userToPrim(user model.User) primitive.M {
	usuario := primitive.M{}

	usuario["_id"], _ = primitive.ObjectIDFromHex(user.ID)
	if user.Name != "" {
		usuario["name"] = user.Name
	}
	if user.Password != "" {
		usuario["password"] = user.Password
	}
	if user.Role != "" {
		usuario["role"] = user.Role
	}
	createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt)
	usuario["createdAt"] = primitive.NewDateTimeFromTime(createdAt)
	updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt)
	usuario["updatedAt"] = primitive.NewDateTimeFromTime(updatedAt)

	return usuario
}

// getLocalUsers retrieves all users from the local database
func getLocalUsers() ([]model.User, error) {
	var users []model.User
	result := database.DBlocal.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
