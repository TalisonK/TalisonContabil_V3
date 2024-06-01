package model

import (
	"context"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCategories() ([]domain.Category, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDBLocal {

		var categories []domain.Category
		if result := database.DBlocal.Find(&categories); result.Error != nil {
			logging.FailedToFindOnDB("All Categories", constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.FoundOnDB("All Categories", constants.LOCAL)
			return categories, nil
		}
	}

	if statusDBCloud {
		result, err := database.DBCloud.Category.Find(context.TODO(), &bson.M{})

		if err != nil {
			logging.FailedToFindOnDB("All Categories", constants.CLOUD, err)
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		var categories []domain.Category

		for result.Next(context.TODO()) {
			var category bson.M
			result.Decode(&category)

			cat := domain.PrimToCategory(category)

			categories = append(categories, cat)
		}

		logging.FoundOnDB("All Categories", constants.CLOUD)
		return categories, nil
	}
	return nil, nil
}

func CreateCategory(category domain.Category) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	category.CreatedAt = util.GetTimeNow()
	category.UpdatedAt = util.GetTimeNow()

	if statusDBCloud {

		pcat := category.ToPrim()

		resultCloud, err := database.DBCloud.Category.InsertOne(context.TODO(), pcat)

		if err != nil {
			logging.FailedToCreateOnDB(category.Name, constants.CLOUD, err)
			return util.GetTagError(http.StatusInternalServerError, err)
		}
		category.ID = resultCloud.InsertedID.(primitive.ObjectID).Hex()
		logging.CreatedOnDB(category.ID, constants.CLOUD)
	}

	if statusDBLocal && category.ID != "" {

		if result := database.DBlocal.Create(&category); result.Error != nil {
			logging.FailedToCreateOnDB(category.ID, constants.LOCAL, result.Error)
			return util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.CreatedOnDB(category.ID, constants.LOCAL)
		}
	}

	return nil
}
func UpdateCategory(category domain.Category) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	baseCategory, err := FindCategoryByID(category.ID)

	if err != nil {
		logging.FailedToFindOnDB(category.ID, constants.LOCAL, err.Inner)
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
			logging.FailedToUpdateOnDB(category.ID, constants.LOCAL, result.Error)
			return util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.UpdatedOnDB(category.ID, constants.LOCAL)
		}
	}

	if statusDBCloud {

		pcat := category.ToPrim()

		id, _ := primitive.ObjectIDFromHex(category.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": pcat}

		_, err := database.DBCloud.Category.UpdateOne(context.Background(), filter, update)

		if err != nil {
			logging.FailedToUpdateOnDB(category.ID, constants.CLOUD, err)
			return util.GetTagError(http.StatusInternalServerError, err)
		}

		logging.UpdatedOnDB(category.ID, constants.CLOUD)
	}

	return nil
}

func DeleteCategory(id string) *util.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logging.FailedToConvertPrimitive(err)
		return util.GetTagError(http.StatusInternalServerError, err)
	}

	if statusDBLocal {
		result := database.DBlocal.Delete(&domain.Category{}, "id=?", id)

		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, constants.LOCAL, result.Error)
			return util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.DeletedOnDB(id, constants.LOCAL)
		}
	}

	if statusDBCloud {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.Category.DeleteOne(context.Background(), filter)

		if err != nil {
			logging.FailedToDeleteOnDB(id, constants.CLOUD, err)
			return util.GetTagError(http.StatusInternalServerError, err)
		}

		logging.DeletedOnDB(id, constants.CLOUD)
	}
	return nil
}

func FindCategoryByID(id string) (*domain.Category, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDBLocal {

		category := domain.Category{}
		result := database.DBlocal.First(&category, "id = ?", id)
		if result.Error != nil {
			logging.FailedToFindOnDB(id, constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			logging.FoundOnDB(id, constants.LOCAL)
			return &category, nil
		}
	}

	if statusDBCloud {
		result := database.DBCloud.Category.FindOne(context.TODO(), bson.M{"_id": id})

		var category domain.Category
		if err := result.Decode(&category); err != nil {
			logging.FailedToFindOnDB(id, constants.CLOUD, err)
			return nil, util.GetTagError(http.StatusInternalServerError, err)
		}

		logging.FoundOnDB(id, constants.CLOUD)
		return &category, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}
