package handler

import (
	"context"

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

			// Get users from local
			result, err = getLocalUsers()
			if err != nil {
				util.LogHandler("Failed to get users from all database.", err, "getUsers")
				return c.Status(500).JSON("Failed to get users from database.")
			}
			util.LogHandler("Users successfully retrieved from local database.", nil, "getUsers")
			return c.Status(200).JSON(result)
		}
		util.LogHandler("Users successfully retrieved from cloud database.", nil, "getUsers")
		return c.Status(200).JSON(result)

	} else {
		// Get users from local
		if database.CheckLocalDB() {
			result, err := getLocalUsers()
			if err != nil {
				util.LogHandler("Failed to get users from local database.", err, "getUsers")
				return c.Status(500).JSON("Failed to get users from database.")
			}
			util.LogHandler("Users successfully retrieved from local database.", nil, "getUsers")
			return c.Status(200).JSON(result)
		} else {
			util.LogHandler("No database connection available.", nil, "getUsers")
			return c.Status(500).JSON("No database connection available.")
		}
	}

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
		usuario.CreateAt = user["createdAt"].(primitive.DateTime).Time().String()
		usuario.UpdateAt = user["updatedAt"].(primitive.DateTime).Time().String()

		result = append(result, usuario)
	}

	if err != nil {
		util.LogHandler("Failed to get users from cloud database.", err, "getUsers")
		return nil, err
	}
	return result, nil
}

func getLocalUsers() ([]model.User, error) {
	colect := database.DBlocal.Find(&model.User{})

	result := []model.User{}

	colect.Scan(&result)

	return result, nil
}
