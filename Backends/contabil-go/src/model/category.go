package model

import (
	"context"
	"fmt"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCategories() ([]domain.Category, error) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		e := logging.NoDatabaseConnection("model.GetCategories")
		return nil, fmt.Errorf(e)
	}

	if statusDBLocal {

		var categories []domain.Category
		if result := database.DBlocal.Find(&categories); result.Error != nil {
			logging.FailedToFindOnDB("All Categories", "Local", result.Error, "model.GetCategories")
			return nil, result.Error
		} else {
			logging.FoundOnDB("All Categories", "Local", "model.GetCategories")
			return categories, nil
		}
	}

	if statusDBCloud {
		result, err := database.DBCloud.Category.Find(context.TODO(), &bson.M{})

		if err != nil {
			logging.FailedToFindOnDB("All Categories", "Cloud", err, "model.GetCategories")
			return nil, err
		}

		var categories []domain.Category

		for result.Next(context.TODO()) {
			var category bson.M
			result.Decode(&category)

			cat := primToCategory(category)

			categories = append(categories, cat)
		}

		logging.FoundOnDB("All Categories", "Cloud", "model.GetCategories")
		return categories, nil
	}
	return nil, nil
}

func CreateCategory(category domain.Category) error {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		e := logging.NoDatabaseConnection("model.CreateCategory")
		return fmt.Errorf(e)
	}

	category.CreatedAt = util.GetTimeNow()
	category.UpdatedAt = util.GetTimeNow()

	if statusDBLocal {

		if result := database.DBlocal.Create(&category); result.Error != nil {
			logging.FailedToCreateOnDB(category.ID, "Local", result.Error, "model.CreateCategory")
			return result.Error
		} else {
			logging.CreatedOnDB(category.ID, "Local", "model.CreateCategory")
		}
	}

	if statusDBCloud {
		_, err := database.DBCloud.Category.InsertOne(context.TODO(), category)

		if err != nil {
			logging.FailedToCreateOnDB(category.ID, "Cloud", err, "model.CreateCategory")
			return err
		}

		logging.CreatedOnDB(category.ID, "Cloud", "model.CreateCategory")
	}

	return nil
}

func UpdateCategory(category domain.Category) error {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		e := logging.NoDatabaseConnection("model.UpdateCategory")
		return fmt.Errorf(e)
	}

	baseCategory, err := FindCategoryByID(category.ID)

	if err != nil {
		logging.FailedToFindOnDB(category.ID, "Local", err, "model.UpdateCategory")
		return err
	}

	if category.Name != "" {
		baseCategory.Name = category.Name
	}

	if category.Description != "" {
		baseCategory.Description = category.Description
	}

	category.CreatedAt = baseCategory.CreatedAt
	category.UpdatedAt = util.GetTimeNow()

	if statusDBLocal {

		if result := database.DBlocal.Save(&baseCategory); result.Error != nil {
			logging.FailedToUpdateOnDB(category.ID, "Local", result.Error, "model.UpdateCategory")
			return result.Error
		} else {
			logging.UpdatedOnDB(category.ID, "Local", "model.UpdateCategory")
		}
	}

	if statusDBCloud {
		_, err := database.DBCloud.Category.ReplaceOne(context.TODO(), bson.M{"_id": category.ID}, category)

		if err != nil {
			logging.FailedToUpdateOnDB(category.ID, "Cloud", err, "model.UpdateCategory")
			return err
		}

		logging.UpdatedOnDB(category.ID, "Cloud", "model.UpdateCategory")
	}

	return nil
}

func DeleteCategory(id string) error {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		e := logging.NoDatabaseConnection("model.DeleteCategory")
		return fmt.Errorf(e)
	}

	if statusDBLocal {
		result := database.DBlocal.Delete(&domain.Category{}, "id=?", id)

		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, "Local", result.Error, "model.DeleteCategory")
			return result.Error
		} else {
			logging.DeletedOnDB(id, "Local", "model.DeleteCategory")
		}
	}

	if statusDBCloud {
		_, err := database.DBCloud.Category.DeleteOne(context.TODO(), bson.M{"_id": id})

		if err != nil {
			logging.FailedToDeleteOnDB(id, "Cloud", err, "model.DeleteCategory")
			return err
		}

		logging.DeletedOnDB(id, "Cloud", "model.DeleteCategory")
	}
	return nil
}

func FindCategoryByID(id string) (*domain.Category, error) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		e := logging.NoDatabaseConnection("model.FindCategoryByID")
		return nil, fmt.Errorf(e)
	}

	if statusDBLocal {

		var category domain.Category
		if result := database.DBlocal.First(&category, id); result.Error != nil {
			logging.FailedToFindOnDB(id, "Local", result.Error, "model.FindCategoryByID")
			return nil, result.Error
		} else {
			logging.FoundOnDB(id, "Local", "model.FindCategoryByID")
			return &category, nil
		}
	}

	if statusDBCloud {
		result := database.DBCloud.Category.FindOne(context.TODO(), bson.M{"_id": id})

		var category domain.Category
		if err := result.Decode(&category); err != nil {
			logging.FailedToFindOnDB(id, "Cloud", err, "model.FindCategoryByID")
			return nil, err
		}

		logging.FoundOnDB(id, "Cloud", "model.FindCategoryByID")
		return &category, nil
	}

	return nil, fmt.Errorf(logging.ErrorOccurred("model.FindCategoryByID"))
}

func primToCategory(prim primitive.M) domain.Category {
	return domain.Category{
		ID:          prim["_id"].(primitive.ObjectID).Hex(),
		Name:        prim["name"].(string),
		Description: prim["description"].(string),
		CreatedAt:   prim["createdAt"].(string),
		UpdatedAt:   prim["updatedAt"].(string),
	}
}

func categoryToPrim(category domain.Category) primitive.M {
	return primitive.M{
		"_id":         category.ID,
		"name":        category.Name,
		"description": category.Description,
		"createdAt":   category.CreatedAt,
		"updatedAt":   category.UpdatedAt,
	}
}
