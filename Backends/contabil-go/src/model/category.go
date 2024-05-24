package model

import (
	"context"
	"fmt"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCategories() ([]domain.Category, error) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		util.LogHandler("No database connection available.", nil, "model.CreateUser")
		return nil, fmt.Errorf("no database connection available")
	}

	if statusDBLocal {

		var categories []domain.Category
		if result := database.DBlocal.Find(&categories); result.Error != nil {
			util.LogHandler("Failed to get categories from local database", result.Error, "model.GetCategories")
			return nil, result.Error
		} else {
			util.LogHandler("Categories retrieved from local database", nil, "model.GetCategories")
			return categories, nil
		}
	}

	if statusDBCloud {
		result, err := database.DBCloud.Category.Find(context.TODO(), &bson.M{})

		if err != nil {
			util.LogHandler("Failed to get categories from cloud database", err, "model.GetCategories")
			return nil, err
		}

		var categories []domain.Category

		for result.Next(context.TODO()) {
			var category bson.M
			result.Decode(&category)

			cat := primToCategory(category)

			categories = append(categories, cat)
		}

		util.LogHandler("Categories retrieved from cloud database", nil, "model.GetCategories")
		return categories, nil
	}
	return nil, nil
}

func primToCategory(prim primitive.M) domain.Category {
	return domain.Category{
		ID:          prim["_id"].(primitive.ObjectID).Hex(),
		Name:        prim["name"].(string),
		Description: prim["description"].(string),
	}
}
