package model

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetExpensesByDate(userId string, startingDate string, endingDate string) ([]domain.Expense, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	expenses := []domain.Expense{}

	if statusDBLocal {

		result := database.DBlocal.Where("user_id = ? AND paid_at between ? AND ?", userId, startingDate, endingDate).Order("created_at DESC").Find(&expenses)

		if result.Error != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Expenses from user %s", userId), constants.LOCAL, result.Error)
			return nil, util.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.FoundOnDB(fmt.Sprintf("Expenses from user %s", userId), constants.LOCAL)
		return expenses, nil

	}

	if statusDBCloud {

		auxSD, _ := time.Parse(time.RFC3339, startingDate)
		sd := primitive.NewDateTimeFromTime(auxSD)

		auxED, _ := time.Parse(time.RFC3339, endingDate)
		ed := primitive.NewDateTimeFromTime(auxED)

		sdBson := bson.M{"$gt": sd, "$lt": ed}
		filter := bson.M{"userID": userId, "paidAt": sdBson}

		opts := options.Find().SetSort(bson.D{{"createdAt", -1}})

		cursor, err := database.DBCloud.Expense.Find(context.Background(), filter, opts)

		if err != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Expenses from user %s", userId), constants.CLOUD, err)
			return nil, util.GetTagError(http.StatusInternalServerError, err)
		}

		for cursor.Next(context.Background()) {
			var raw bson.M
			cursor.Decode(raw)

			expenses = append(expenses, domain.PrimToExpense(raw))
		}

		return expenses, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}
