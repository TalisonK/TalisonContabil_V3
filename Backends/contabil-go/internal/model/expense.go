package model

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetExpensesByDate(userId string, startingDate string, endingDate string, statusDBLocal bool, statusDBCloud bool) ([]domain.ExpenseDTO, *tagError.TagError) {

	expenses := []domain.Expense{}

	if statusDBLocal {

		result := database.DBlocal.
			Where("user_id = ? AND paid_at between ? AND ?", userId, startingDate, endingDate).
			Joins("Category", "expense.category_id = categories.id").
			Order("created_at DESC").
			Find(&expenses)

		if result.Error != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Expenses from user %s", userId), constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.FoundOnDB(fmt.Sprintf("Expenses from user %s", userId), constants.LOCAL)

		dtos := ExpenseGetCategoryName(expenses, statusDBLocal, statusDBCloud)

		return dtos, nil
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
			return nil, tagError.GetTagError(http.StatusInternalServerError, err)
		}

		for cursor.Next(context.Background()) {
			var raw bson.M

			cursor.Decode(raw)

			expenses = append(expenses, domain.PrimToExpense(raw))
		}

		dtos := ExpenseGetCategoryName(expenses, statusDBLocal, statusDBCloud)

		return dtos, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func CreateExpense() {}

func ExpenseGetCategoryName(expenses []domain.Expense, statusDBLocal bool, statusDBCloud bool) []domain.ExpenseDTO {
	expensesDto := make([]domain.ExpenseDTO, len(expenses))

	errors := make(chan *tagError.TagError, len(expenses))

	var wg sync.WaitGroup
	for i, exp := range expenses {
		wg.Add(1)

		go func(i int, exp domain.Expense) {
			defer wg.Done()

			cat, tagErr := FindCategoryByID(exp.CategoryID, statusDBLocal, statusDBCloud)

			if tagErr != nil {
				logging.FailedToFindOnDB(fmt.Sprintf("Category from user %s", exp.UserID), constants.LOCAL, tagErr.Inner)
				errors <- tagErr
				return
			}

			expDto := exp.ToDTO()
			expDto.CategoryName = cat.Name

			expensesDto[i] = expDto

		}(i, exp)
	}

	wg.Wait()

	return expensesDto
}
