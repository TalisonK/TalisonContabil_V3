package model

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCreditCardsByUser(userId string) []domain.CreditCard {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil
	}

	if statusDbLocal {

		var result []domain.CreditCard

		database.DBlocal.Where("user_id = ?", userId).Find(&result)

		return result
	}

	if statusDbCloud {

		cursor, err := database.DBCloud.CreditCard.Find(context.Background(), primitive.M{"userID": userId})

		if err != nil {
			return nil
		}

		var result []domain.CreditCard

		for cursor.Next(context.Background()) {
			var creditCard domain.CreditCard
			cursor.Decode(&creditCard)
			result = append(result, creditCard)
		}

		return result
	}
	return nil
}

func CreateCreditCard(entry domain.CreditCard) (*domain.CreditCard, *tagError.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	entry.CreatedAt = timeHandler.GetTimeNow()
	entry.UpdatedAt = timeHandler.GetTimeNow()

	if statusDbCloud {

		prim := entry.ToPrim()

		inserted, err := database.DBCloud.CreditCard.InsertOne(context.Background(), prim)

		if err != nil {
			return nil, tagError.GetTagError(http.StatusInternalServerError, fmt.Errorf(logging.FailedToCreateOnDB(entry.Flag, constants.CLOUD, err)))
		}

		entry.ID = inserted.InsertedID.(primitive.ObjectID).Hex()

		logging.CreatedOnDB(entry.Flag, constants.CLOUD)
	}

	if statusDbLocal {

		result := database.DBlocal.Create(&entry)

		if result.Error != nil {
			return nil, tagError.GetTagError(http.StatusInternalServerError, fmt.Errorf(logging.FailedToCreateOnDB(entry.Flag, constants.LOCAL, result.Error)))
		}

		logging.CreatedOnDB(entry.Flag, constants.LOCAL)

		return &entry, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

}

func GetEndDay(id string) int {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return -1
	}

	if statusDbLocal {

		var result domain.CreditCard

		database.DBlocal.Where("id = ?", id).First(&result)

		return result.ExpiresAt
	}

	if statusDbCloud {

		prim, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return -1
		}

		result := database.DBCloud.CreditCard.FindOne(context.Background(), primitive.M{"_id": prim})

		if result.Err() != nil {
			return -1
		}

		var creditCard domain.CreditCard

		result.Decode(&creditCard)

		return creditCard.ExpiresAt
	}

	return -1
}
