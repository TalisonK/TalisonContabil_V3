package model

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetIncomes() ([]domain.Income, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.GetIncomes"))
	}

	if !statusDBLocal {
		incomes := []domain.Income{}
		err := database.DBlocal.Find(&incomes).Error

		if err != nil {
			e := logging.FailedToFindOnDB("All Incomes", "Incomes", nil, "model.GetIncomes")
			return nil, util.GetTagError(http.StatusInternalServerError, fmt.Errorf(e))
		}

		logging.FoundOnDB("All Incomes", "Local", "model.GetIncomes")
		return incomes, nil
	}

	if statusDBCloud {

		filter := bson.M{}

		cursor, err := database.DBCloud.Income.Find(context.Background(), filter)

		if err != nil {
			e := logging.FailedToFindOnDB("All Incomes", "Incomes", err, "model.GetIncomes")
			return nil, util.GetTagError(http.StatusInternalServerError, fmt.Errorf(e))
		}
		incomes := []domain.Income{}

		for cursor.Next(context.Background()) {
			var income primitive.M
			cursor.Decode(&income)

			incomes = append(incomes, primToIncome(income))
		}

		logging.FoundOnDB("All Incomes", "Cloud", "model.GetIncomes")

		return incomes, nil

	}
	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.GetIncomes"))
}

func primToIncome(income primitive.M) domain.Income {
	newIncome := domain.Income{}

	newIncome.ID = income["_id"].(primitive.ObjectID).Hex()
	newIncome.Value = income["value"].(float64)
	newIncome.Description = income["description"].(string)
	newIncome.ReceivedAt = income["receivedAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	user := income["user"].(primitive.M)
	newIncome.UserID = user["_id"].(primitive.ObjectID).Hex()
	newIncome.CreatedAt = income["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)

	if income["updatedAt"] != nil {
		newIncome.UpdatedAt = income["updatedAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	}

	return newIncome
}
