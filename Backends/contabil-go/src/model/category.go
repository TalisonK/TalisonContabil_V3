package model

import (
	"context"
	"net/http"
	"time"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCategories() ([]domain.Category, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.GetCategory"))
	}

	if statusDBLocal {

		var categories []domain.Category
		if result := database.DBlocal.Find(&categories); result.Error != nil {
			logging.FailedToFindOnDB("All Categories", "Local", result.Error, "model.GetCategories")
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.FoundOnDB("All Categories", "Local", "model.GetCategories")
			return categories, nil
		}
	}

	if statusDBCloud {
		result, err := database.DBCloud.Category.Find(context.TODO(), &bson.M{})

		if err != nil {
			logging.FailedToFindOnDB("All Categories", "Cloud", err, "model.GetCategories")
			return nil, util.GetTagError(http.StatusBadRequest, err)
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

func CreateCategory(category domain.Category) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.CreateCategory"))
	}

	category.CreatedAt = util.GetTimeNow()
	category.UpdatedAt = util.GetTimeNow()

	if statusDBCloud {

		pcat := categoryToPrim(category)

		resultCloud, err := database.DBCloud.Category.InsertOne(context.TODO(), pcat)

		if err != nil {
			logging.FailedToCreateOnDB(category.ID, "Cloud", err, "model.CreateCategory")
			return util.GetTagError(http.StatusInternalServerError, err)
		}
		category.ID = resultCloud.InsertedID.(primitive.ObjectID).Hex()
		logging.CreatedOnDB(category.ID, "Cloud", "model.CreateCategory")
	}

	if statusDBLocal {

		if result := database.DBlocal.Create(&category); result.Error != nil {
			logging.FailedToCreateOnDB(category.ID, "Local", result.Error, "model.CreateCategory")
			return util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.CreatedOnDB(category.ID, "Local", "model.CreateCategory")
		}
	}

	return nil
}
func UpdateCategory(category domain.Category) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.UpdateCategory"))
	}

	baseCategory, err := FindCategoryByID(category.ID)

	if err != nil {
		logging.FailedToFindOnDB(category.ID, "Local", err.Inner, "model.UpdateCategory")
		return err
	}

	if category.Name != "" {
		baseCategory.Name = category.Name
	}

	if category.Description != "" {
		baseCategory.Description = category.Description
	}

	category.CreatedAt = baseCategory.CreatedAt

	category = *baseCategory
	category.UpdatedAt = util.GetTimeNow()

	if statusDBLocal {

		if result := database.DBlocal.Save(&category); result.Error != nil {
			logging.FailedToUpdateOnDB(category.ID, "Local", result.Error, "model.UpdateCategory")
			return util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.UpdatedOnDB(category.ID, "Local", "model.UpdateCategory")
		}
	}

	if statusDBCloud {

		pcat := categoryToPrim(category)

		id, _ := primitive.ObjectIDFromHex(category.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": pcat}

		_, err := database.DBCloud.Category.UpdateOne(context.Background(), filter, update)

		if err != nil {
			logging.FailedToUpdateOnDB(category.ID, "Cloud", err, "model.UpdateCategory")
			return util.GetTagError(http.StatusInternalServerError, err)
		}

		logging.UpdatedOnDB(category.ID, "Cloud", "model.UpdateCategory")
	}

	return nil
}

func DeleteCategory(id string) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.DeleteCategory"))
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logging.FailedToConvertPrimitive(err, "model.DeleteUser")
		return util.GetTagError(http.StatusInternalServerError, err)
	}

	if statusDBLocal {
		result := database.DBlocal.Delete(&domain.Category{}, "id=?", id)

		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, "Local", result.Error, "model.DeleteCategory")
			return util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.DeletedOnDB(id, "Local", "model.DeleteCategory")
		}
	}

	if statusDBCloud {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.Category.DeleteOne(context.Background(), filter)

		if err != nil {
			logging.FailedToDeleteOnDB(id, "Cloud", err, "model.DeleteCategory")
			return util.GetTagError(http.StatusInternalServerError, err)
		}

		logging.DeletedOnDB(id, "Cloud", "model.DeleteCategory")
	}
	return nil
}

func FindCategoryByID(id string) (*domain.Category, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.FindCategoryByID"))
	}

	if statusDBLocal {

		category := domain.Category{}
		result := database.DBlocal.First(&category, "id = ?", id)
		if result.Error != nil {
			logging.FailedToFindOnDB(id, "Local", result.Error, "model.FindCategoryByID")
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
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
			return nil, util.GetTagError(http.StatusInternalServerError, err)
		}

		logging.FoundOnDB(id, "Cloud", "model.FindCategoryByID")
		return &category, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.FindCategoryByID"))
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

	pcat := primitive.M{}

	if category.ID != "" {
		id, _ := primitive.ObjectIDFromHex(category.ID)
		pcat["_id"] = id
	}

	pcat["name"] = category.Name
	pcat["description"] = category.Description

	createdAt, _ := time.Parse(time.RFC3339, category.CreatedAt)
	pcat["createdAt"] = primitive.NewDateTimeFromTime(createdAt)

	updatedAt, _ := time.Parse(time.RFC3339, category.UpdatedAt)
	pcat["updatedAt"] = primitive.NewDateTimeFromTime(updatedAt)

	return pcat
}
