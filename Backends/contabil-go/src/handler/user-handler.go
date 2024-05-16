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
		user.CreatedAt = now.Local().Format("2023-10-23T20:49:22.723+00:00")
		user.UpdatedAt = now.Local().Format("2023-10-23T20:49:22.723+00:00")
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

func getCloudUsers() ([]model.User, error) {
	colect, err := database.DBCloud.User.Find(context.TODO(), &bson.D{})

	result := []model.User{}

	//colect.All(context.Background(), &result)

	for colect.Next(context.Background()) {
		var user bson.M
		colect.Decode(&user)

		var usuario model.User

		usuario.ID = user["_id"].(primitive.ObjectID).Hex()
		usuario.Name = user["name"].(string)
		usuario.Password = user["password"].(string)
		usuario.Role = user["role"].(string)
		usuario.CreatedAt = user["createdAt"].(primitive.DateTime).Time().String()
		usuario.UpdatedAt = user["updatedAt"].(primitive.DateTime).Time().String()

		result = append(result, usuario)
	}

	if err != nil {
		util.LogHandler("Failed to get users from cloud database.", err, "getUsers")
		return nil, err
	}
	return result, nil
}

func getLocalUsers() ([]model.User, error) {
	var users []model.User
	result := database.DBlocal.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
