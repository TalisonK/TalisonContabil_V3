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
	"github.com/TalisonK/TalisonContabil/pkg/mathPlus"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TotalRanger(ctx context.Context, cancel func(), userID string, originMonth string, originYear int) ([]domain.IncomevsExpense, *tagError.TagError) {

	// creating arrays with pre-loaded size
	grathData := make([]domain.IncomevsExpense, 13)
	errors := make(chan *tagError.TagError, 13)

	// Get starting date
	month, year := timeHandler.MonthSubtractorByJump(originMonth, originYear, 5)

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	var wg sync.WaitGroup
	for i := 0; i < 13; i++ {
		wg.Add(1)

		go func(i int, month string, year int, statusDBLocal bool, statusDBCloud bool) {

			select {
			case <-ctx.Done():
				// Context was cancelled, return to prevent further processing
				logging.ContextAlreadyClosed()
				return

			default:
				// Context was not cancelled, continue processing
				defer wg.Done()
				actual, err := mountInvsEx(userID, month, year, statusDBLocal, statusDBCloud)
				if err != nil {
					logging.FailedToCreateOnDB(fmt.Sprintf("IncomeVSExpense for %s/%d", month, year), constants.ALL, err.Inner)
					errors <- err
					cancel()
					return
				}
				grathData[i] = actual
			}

		}(i, month, year, statusDBLocal, statusDBCloud)
		month, year = timeHandler.MonthAdder(month, year)
	}

	wg.Wait()
	close(errors)

	if len(errors) > 0 {
		return nil, <-errors
	}

	return grathData, nil
}

func mountInvsEx(userID string, month string, year int, statusDBLocal bool, statusDBCloud bool) (domain.IncomevsExpense, *tagError.TagError) {
	actual := domain.IncomevsExpense{}

	var income *domain.Total
	var expense *domain.Total

	income, tagErr := findTotalByMonthAndYear(month, year, constants.INCOME, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("%ss for user %s", constants.INCOME, userID), constants.ALL, tagErr.Inner)
	}

	if income.ID == "" {
		income, tagErr = CreateUpdateTotal(userID, month, year, constants.INCOME, statusDBLocal, statusDBCloud)

		if tagErr != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("%ss for user %s", constants.INCOME, userID), constants.ALL, tagErr.Inner)
			return domain.IncomevsExpense{}, tagErr
		}
	}

	expense, tagErr = findTotalByMonthAndYear(month, year, constants.EXPENSE, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("%ss for user %s", constants.EXPENSE, userID), constants.ALL, tagErr.Inner)
		return domain.IncomevsExpense{}, tagErr
	}

	if expense.ID == "" {
		expense, tagErr = CreateUpdateTotal(userID, month, year, constants.EXPENSE, statusDBLocal, statusDBCloud)

		if tagErr != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("%ss for user %s", constants.EXPENSE, userID), constants.ALL, tagErr.Inner)
			return domain.IncomevsExpense{}, tagErr
		}
	}

	actual.Income = income.TotalValue
	actual.Expense = expense.TotalValue
	actual.Month = month
	actual.Year = year

	return actual, nil
}

func CreateUpdateTotal(userId string, month string, year int, totalType string, statusDBLocal bool, statusDBCloud bool) (*domain.Total, *tagError.TagError) {

	// check if the month and year are valid

	if month == "" || year == 0 {
		return nil, tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
	}

	// get the first and last day of the month
	startingDate, endingDate := timeHandler.GetFirstAndLastDayOfMonth(month, year)

	// fetch the incomes or expenses from the database
	var activities []domain.Activity
	var tagErr *tagError.TagError

	if totalType == constants.INCOME {
		activities, tagErr = fetchIncomesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)
	} else {
		activities, tagErr = fetchExpensesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)
	}

	if tagErr != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("%ss for user %s", totalType, userId), constants.INCOME, tagErr.Inner)
		return nil, tagErr
	}

	// mount the total
	total := mountTotal(month, year, userId, totalType, activities)

	// check if the total already exists
	old, tagErr := findTotalByMonthAndYear(month, year, totalType, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Totals for user %s", userId), "Total", tagErr.Inner)
		return nil, tagErr
	}

	// create or update the total in the database
	if old.ID != "" {
		total.ID = old.ID
		total.CreatedAt = old.CreatedAt
	}

	if statusDBCloud {

		if old.ID == "" {
			total, tagErr = createTotalInDB(total)
		} else {
			total, tagErr = updateTotalInDB(total)
		}

		if tagErr != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", month, year), constants.CLOUD, tagErr.Inner)
			return nil, tagErr
		}

	}

	if statusDBLocal {
		result := database.DBlocal.Save(&total)

		if result.Error != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Total for %s from %s/%d", totalType, month, year), constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusInternalServerError, result.Error)
		}

		logging.CreatedOnDB(total.ID, constants.LOCAL)
		return &total, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func resumeBalance(actual float64, pass float64) float64 {

	if actual == 0 || pass == 0 {
		return 0.0
	}

	x := (100 * pass) / actual

	return float64(mathPlus.ToFixed(100-x, 2))
}

func fetchIncomesByDate(userId string, startingDate string, endingDate string, statusDBLocal bool, statusDBCloud bool) ([]domain.Activity, *tagError.TagError) {

	var activities []domain.Activity

	incomes, tagError := GetIncomesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Incomes for user %s", userId), constants.INCOME, tagError.Inner)
		return nil, tagError
	}

	for _, income := range incomes {
		activities = append(activities, income.ToActivity())
	}

	return activities, nil

}

func fetchExpensesByDate(userId string, startingDate string, endingDate string, statusDBLocal bool, statusDBCloud bool) ([]domain.Activity, *tagError.TagError) {

	var activities []domain.Activity

	expenses, tagError := GetExpensesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)

	if tagError != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Expenses for user %s", userId), constants.INCOME, tagError.Inner)
		return nil, tagError
	}

	for _, expenses := range expenses {
		activities = append(activities, expenses.ToActivity())
	}

	return activities, nil

}

func Timeline(ctx context.Context, cancel func(), errChan chan *tagError.TagError, userId string, month string, year int) ([]domain.Activity, *tagError.TagError) {

	var incomes []domain.Activity
	var expenses []domain.Activity

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	startingDate, endingDate := timeHandler.GetFirstAndLastDayOfMonth(month, year)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		defer wg.Done()
		select {
		case <-ctx.Done():
			// Context was cancelled, return to prevent further processing
			logging.ContextAlreadyClosed()
			return
		default:
			result, tagError := fetchIncomesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)

			if tagError != nil {
				logging.FailedToFindOnDB(fmt.Sprintf("Incomes for user %s", userId), constants.INCOME, tagError.Inner)
				errChan <- tagError
				cancel()
				return
			}

			logging.FoundOnDB(fmt.Sprintf("Incomes for user %s", userId), constants.INCOME)
			incomes = result
		}

	}()

	go func() {

		defer wg.Done()
		select {
		case <-ctx.Done():
			// Context was cancelled, return to prevent further processing
			logging.ContextAlreadyClosed()
			return
		default:
			result, tagError := fetchExpensesByDate(userId, startingDate, endingDate, statusDBLocal, statusDBCloud)

			if tagError != nil {
				logging.FailedToFindOnDB(fmt.Sprintf("Expenses for user %s", userId), constants.INCOME, tagError.Inner)
				errChan <- tagError
				cancel()
				return
			}

			logging.FoundOnDB(fmt.Sprintf("Expenses for user %s", userId), constants.INCOME)
			expenses = result
		}

	}()

	wg.Wait()

	// In case one of the arrays are empty, just return the opposite

	inLen := len(incomes)
	exLen := len(expenses)

	if inLen == 0 {
		return expenses, nil
	}
	if exLen == 0 {
		return incomes, nil
	}

	size := inLen + exLen

	exIndex := 0
	inIndex := 0

	timeline := []domain.Activity{}

	// Sort and stacking activities

	for i := 0; i < size; i++ {

		if exIndex == exLen {
			timeline = append(timeline, incomes...)
			break
		}

		if inIndex == inLen {
			timeline = append(timeline, expenses[exIndex:]...)
			break
		}

		expense, _ := time.Parse(time.RFC3339, expenses[exIndex].ActivityDate)
		income, _ := time.Parse(time.RFC3339, incomes[inIndex].ActivityDate)

		switch income.Compare(expense) {
		case -1:
			timeline = append(timeline, expenses[exIndex])
			exIndex++

		default:
			timeline = append(timeline, incomes[inIndex])
			inIndex++
		}
	}

	return timeline, nil

}

func Resume(ive []domain.IncomevsExpense) (map[string]domain.Resume, *tagError.TagError) {

	resumes := map[string]domain.Resume{}

	pass := ive[5]
	actual := ive[6]

	resumes["income"] = domain.Resume{
		Actual:  actual.Income,
		Pass:    pass.Income,
		Balance: resumeBalance(actual.Income, pass.Income),
	}

	resumes["expense"] = domain.Resume{
		Actual:  actual.Expense,
		Pass:    pass.Expense,
		Balance: resumeBalance(actual.Expense, pass.Expense),
	}

	resumes["balance"] = domain.Resume{
		Actual:  mathPlus.ToFixed(actual.Income-actual.Expense, 2),
		Pass:    mathPlus.ToFixed(pass.Income-pass.Expense, 2),
		Balance: resumeBalance((actual.Income - actual.Expense), (pass.Income - pass.Expense)),
	}

	return resumes, nil

}

// createTotalInDB creates the total in the database
func createTotalInDB(total domain.Total) (domain.Total, *tagError.TagError) {

	raw := total.ToPrim()

	inserted, err := database.DBCloud.Total.InsertOne(context.Background(), raw)
	if err != nil {
		logging.FailedToCreateOnDB(fmt.Sprintf("Total for income from %s/%d", total.Month, total.Year), constants.CLOUD, err)
		return domain.Total{}, tagError.GetTagError(http.StatusInternalServerError, err)
	}

	total.ID = inserted.InsertedID.(primitive.ObjectID).Hex()
	logging.CreatedOnDB(total.ID, constants.CLOUD)
	return total, nil
}

// updateTotalInDB updates the total in the database
func updateTotalInDB(total domain.Total) (domain.Total, *tagError.TagError) {

	total.UpdatedAt = timeHandler.GetTimeNow()

	filter := bson.M{"_id": total.ID}

	parser := bson.M{"$set": total.ToPrim()}

	_, err := database.DBCloud.Total.UpdateOne(context.Background(), filter, parser)

	if err != nil {
		logging.FailedToUpdateOnDB(fmt.Sprintf("Total for income from %s/%d", total.Month, total.Year), constants.CLOUD, err)
		return domain.Total{}, tagError.GetTagError(http.StatusInternalServerError, err)
	}

	logging.UpdatedOnDB(total.ID, constants.CLOUD)
	return total, nil
}

func findTotalByMonthAndYear(month string, year int, totalType string, statusDBLocal bool, statusDBCloud bool) (*domain.Total, *tagError.TagError) {

	if statusDBLocal {

		var total domain.Total

		result := database.DBlocal.Where("month = ? AND year = ? AND type = ?", month, year, totalType).Find(&total)

		if result.Error != nil {
			logging.FailedToFindOnDB("Income Total", constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.FoundOnDB(fmt.Sprintf("%ss Total from %s/%d", totalType, month, year), constants.LOCAL)
		return &total, nil
	}

	if statusDBCloud {

		filter := bson.M{"month": month, "year": year, "type": totalType}

		cursor := database.DBCloud.Total.FindOne(context.Background(), filter)

		var raw bson.M

		err := cursor.Decode(&raw)

		if err != nil {
			logging.FailedToFindOnDB("Income Total", constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		total := domain.PrimToTotal(raw)

		logging.FoundOnDB(fmt.Sprintf("%ss Total from %s/%d", totalType, month, year), constants.LOCAL)
		return &total, nil

	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

}

func mountTotal(month string, year int, userId string, totalType string, activities []domain.Activity) domain.Total {
	total := domain.Total{}

	total.CreatedAt = timeHandler.GetTimeNow()
	total.UpdatedAt = timeHandler.GetTimeNow()
	total.Type = totalType
	total.Month = month
	total.Year = year
	total.UserID = userId

	for _, activity := range activities {
		total.TotalValue += activity.Value
	}

	total.TotalValue = mathPlus.ToFixed(total.TotalValue, 2)

	return total
}
