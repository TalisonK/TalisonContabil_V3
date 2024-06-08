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
	"github.com/TalisonK/TalisonContabil/pkg/mathplus"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
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

func ExpenseByMethod(ctx context.Context, cancel func(), errChan chan *tagError.TagError, userId string, month string, year int) (map[string]float64, *tagError.TagError) {
	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	// get the first and last day of the month
	startingDate, endingDate := timeHandler.GetFirstAndLastDayOfMonth(month, year)

	expenses, tagErr := GetExpensesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FoundOnDB(fmt.Sprintf("Expenses for user %s", userId), "User")
		return nil, tagErr
	}

	methods := make(map[string]float64)

	for _, expense := range expenses {
		methods[expense.PaymentMethod] += expense.Value
	}

	return methods, nil

}

func ExpenseByCategory(ctx context.Context, cancel func(), errChan chan *tagError.TagError, userId string, month string, year int) (map[string]float64, *tagError.TagError) {

	expVSCat := make(map[string]float64, len(database.CacheDatabase.Categories))

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	// get the first and last day of the month
	startingDate, endingDate := timeHandler.GetFirstAndLastDayOfMonth(month, year)

	var wg sync.WaitGroup

	for id, cat := range database.CacheDatabase.Categories {

		wg.Add(1)

		select {
		case <-ctx.Done():
			// Context was cancelled, return to prevent further processing
			logging.ContextAlreadyClosed()
			return nil, nil

		default:
			go func(categoryId string, cat domain.Category, statusDBLocal bool, statusDBCloud bool, startingDate string, endingDate string) {

				defer wg.Done()
				expenses, tagErr := findExpenseByCategoryId(categoryId, statusDBLocal, statusDBCloud, startingDate, endingDate)

				if tagErr != nil {
					logging.GenericError(fmt.Sprintf("Failed to create expenses vs Category for category %s", categoryId), tagErr.Inner)
					errChan <- tagErr
					cancel()
					return
				}

				value := 0.0

				for _, expense := range expenses {
					value += expense.Value
				}
				if value > 1 {
					expVSCat[cat.Name] = mathplus.ToFixed(value, 2)
				}

			}(id, cat, statusDBLocal, statusDBCloud, startingDate, endingDate)
		}
	}

	wg.Wait()

	return expVSCat, nil
}

func findExpenseByCategoryId(categoryId string, statusDBLocal bool, statusDBCloud bool, startingDate string, endingDate string) ([]domain.Expense, *tagError.TagError) {

	var expenses []domain.Expense

	if statusDBLocal {
		result := database.DBlocal.Where("category_id = ? AND paid_at between ? AND ?", categoryId, startingDate, endingDate).Find(&expenses)

		if result.Error != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Expenses for category %s", categoryId), constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.FoundOnDB(fmt.Sprintf("Expenses for category %s", categoryId), constants.LOCAL)
		return expenses, nil
	}

	if statusDBCloud {

		auxSD, _ := time.Parse(time.RFC3339, startingDate)
		sd := primitive.NewDateTimeFromTime(auxSD)

		auxED, _ := time.Parse(time.RFC3339, endingDate)
		ed := primitive.NewDateTimeFromTime(auxED)

		sdBson := bson.M{"$gt": sd, "$lt": ed}
		filter := bson.M{"categoryId": categoryId, "paidAt": sdBson}

		cursor, err := database.DBCloud.Expense.Find(context.Background(), filter)

		if err != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Expenses for category %s", categoryId), constants.LOCAL, err)
			return nil, tagError.GetTagError(http.StatusInternalServerError, err)
		}

		expenses := []domain.Expense{}

		for cursor.Next(context.Background()) {
			var aux bson.M

			cursor.Decode(aux)

			expenses = append(expenses, domain.PrimToExpense(aux))
		}

		logging.FoundOnDB(fmt.Sprintf("Expenses for category %s", categoryId), constants.LOCAL)
		return expenses, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

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
