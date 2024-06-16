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

func CreateExpenseHandler(expense domain.ExpenseDTO) ([]string, *tagError.TagError) {

	if expense.CategoryName == "" || expense.PaidAt == "" || expense.PaymentMethod == "" || expense.Value == 0 || expense.Description == "" {
		return nil, tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
	}

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	// validar se a categoria existe

	category, tagErr := FindCategoryByName(expense.CategoryName, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(fmt.Sprintf("Category %s", expense.CategoryName), "Category", tagErr.Inner)
		return nil, tagErr
	}

	if category.ID == "" {
		return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.FailedToFindOnDB(expense.CategoryName, "Category", tagErr.Inner)))
	}

	expense.CategoryID = category.ID

	// validar se o método de pagamento é válido

	cond := true

	for _, method := range constants.GetMethods() {
		if method == expense.PaymentMethod {
			cond = false
			break
		}
	}

	if cond {
		return nil, tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
	}

	// validar se é repetido

	if checkDuplicatedExpense(expense, statusDBLocal, statusDBCloud) {
		return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.DuplicatedEntry(expense.Description)))
	}

	expenses := []string{}

	if expense.PaymentMethod == "CREDIT_CARD" {
		if expense.TotalParcel == 0 || expense.ActualParcel == 0 || expense.TotalParcel < expense.ActualParcel {
			return nil, tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
		}

		var wg sync.WaitGroup
		for i := expense.ActualParcel; i <= expense.TotalParcel; i++ {
			wg.Add(1)

			go func(i int32) {
				defer wg.Done()

				expenseAux := expense
				expenseAux.ActualParcel = i
				expenseAux.ID = ""

				id, tagErr := CreateExpense(expenseAux, statusDBLocal, statusDBCloud)

				if tagErr != nil {
					logging.GenericError(fmt.Sprintf("Failed to create expense %d/%d", i, expense.TotalParcel), tagErr.Inner)
					return
				}
				logging.CreatedOnDB(fmt.Sprintf("Expense from date %d/%d", i, expense.TotalParcel), "Expenses")
				expenses = append(expenses, id.ID)
			}(i)
		}

		wg.Wait()

		month, year := timeHandler.DateBreaker(expense.PaidAt)

		CreateUpdateTotal(expense.UserID, month, year, constants.EXPENSE, statusDBLocal, statusDBCloud)

		logging.GenericSuccess(fmt.Sprintf("Expenses %d/%d created successfully", expense.ActualParcel, expense.TotalParcel))

		return expenses, nil

	}

	id, tagErr := CreateExpense(expense, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.GenericError(fmt.Sprintf("Failed to create expense %s", expense.Description), tagErr.Inner)
		return nil, tagErr
	}

	// Debit card or cash

	logging.CreatedOnDB(fmt.Sprintf("Expense %s", expense.Description), "Expenses")

	expenses = append(expenses, id.ID)

	return expenses, nil
}

func UpdateExpenseHandler(expense domain.ExpenseDTO) ([]string, *tagError.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if expense.ID == "" {
		return nil, tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
	}

	expense = makeExpenseParser(expense)

	expensesToUpdate, tagErr := makeExpensesToUpdate(expense, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(expense.Description, constants.LOCAL, tagErr.Inner)
		return nil, tagErr
	}

	expenses := []string{}

	var wg sync.WaitGroup
	var errors = make(chan *tagError.TagError, len(expensesToUpdate))

	for _, exp := range expensesToUpdate {
		wg.Add(1)

		go func(exp domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) {
			defer wg.Done()

			id, tagErr := UpdateExpense(exp, statusDBLocal, statusDBCloud)

			if tagErr != nil {
				logging.GenericError(fmt.Sprintf("Failed to update expense %s", exp.Description), tagErr.Inner)
				errors <- tagErr
				return
			}

			logging.UpdatedOnDB(fmt.Sprintf("Expense %s", exp.Description), constants.EXPENSE)
			expenses = append(expenses, id.ID)
		}(exp, statusDBLocal, statusDBCloud)

	}

	wg.Wait()

	if len(errors) > 0 {
		return nil, <-errors
	}

	month, year := timeHandler.DateBreaker(expense.PaidAt)

	CreateUpdateTotal(expense.UserID, month, year, constants.EXPENSE, statusDBLocal, statusDBCloud)

	logging.GenericSuccess(fmt.Sprintf("Expenses %s updated successfully", expense.Description))

	return expenses, nil

}

func DeleteExpenseHandler(id string) *tagError.TagError {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if id == "" {
		return tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
	}

	expense := domain.ExpenseDTO{}
	expense.ID = id

	expensesToUpdate, tagErr := makeExpensesToUpdate(expense, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(expense.Description, constants.LOCAL, tagErr.Inner)
		return tagErr
	}

	var wg sync.WaitGroup
	var errors = make(chan *tagError.TagError, 1)

	for _, exp := range expensesToUpdate {
		wg.Add(1)

		go func(exp domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) {
			defer wg.Done()

			tagErr := DeleteExpense(exp, statusDBLocal, statusDBCloud)

			if tagErr != nil {
				logging.GenericError(fmt.Sprintf("Failed to delete expense %s", exp.Description), tagErr.Inner)
				errors <- tagErr
				return
			}

			logging.DeletedOnDB(fmt.Sprintf("Expense %s", exp.Description), constants.EXPENSE)
		}(exp, statusDBLocal, statusDBCloud)
	}

	wg.Wait()

	if len(errors) > 0 {
		return <-errors
	}

	month, year := timeHandler.DateBreaker(expense.PaidAt)

	CreateUpdateTotal(expense.UserID, month, year, constants.EXPENSE, statusDBLocal, statusDBCloud)

	return nil
}

func UpdateExpense(expense domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) (*domain.ExpenseDTO, *tagError.TagError) {

	expenseEntity := expense.ToEntity()

	expenseEntity.UpdatedAt = timeHandler.GetTimeNow()

	if statusDBLocal {
		result := database.DBlocal.Save(&expenseEntity)

		if result.Error != nil {
			logging.FailedToUpdateOnDB(expenseEntity.ID, constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.UpdatedOnDB(expenseEntity.ID, constants.LOCAL)
	}

	if statusDBCloud {
		objId, err := primitive.ObjectIDFromHex(expenseEntity.ID)

		if err != nil {
			logging.FailedToUpdateOnDB(expenseEntity.ID, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		filter := bson.M{"_id": objId}
		update := bson.M{"$set": expenseEntity.ToPrim()}

		_, err = database.DBCloud.Expense.UpdateOne(context.Background(), filter, update)

		if err != nil {
			logging.FailedToUpdateOnDB(expenseEntity.ID, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

}

func CreateExpense(expenseDto domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) (*domain.ExpenseDTO, *tagError.TagError) {

	/*
		userId
		categoryname
		paidAt
		paymentMethod
		actualParcel
		totalParcel
		value
	*/

	expense := expenseDto.ToEntity()

	expense.CreatedAt = timeHandler.GetTimeNow()
	expense.UpdatedAt = timeHandler.GetTimeNow()

	// pegar o id da categoria

	if statusDBCloud {

		inserted, err := database.DBCloud.Expense.InsertOne(context.Background(), expense.ToPrim())

		if err != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Expense %s", expense.Description), "Cloud", err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		expense.ID = inserted.InsertedID.(primitive.ObjectID).Hex()
		logging.CreatedOnDB(fmt.Sprintf("Expense %s", expense.Description), "Cloud")
	}

	if statusDBLocal {

		result := database.DBlocal.Create(&expense)

		if result.Error != nil {
			logging.FailedToCreateOnDB(fmt.Sprintf("Expense %s", expense.Description), "Local", result.Error)
			return nil, tagError.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.CreatedOnDB(fmt.Sprintf("Expense %s", expense.Description), "Local")

		dto := expense.ToDTO()

		return &dto, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func DeleteExpense(expense domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) *tagError.TagError {

	if statusDBLocal {
		result := database.DBlocal.Delete(&domain.Expense{}, "id = ?", expense.ID)

		if result.Error != nil {
			logging.FailedToDeleteOnDB(expense.ID, constants.LOCAL, result.Error)
			return tagError.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.DeletedOnDB(expense.ID, constants.LOCAL)
	}

	if statusDBCloud {
		objId, err := primitive.ObjectIDFromHex(expense.ID)

		if err != nil {
			logging.FailedToDeleteOnDB(expense.ID, constants.CLOUD, err)
			return tagError.GetTagError(http.StatusBadRequest, err)
		}

		filter := bson.M{"_id": objId}

		_, err = database.DBCloud.Expense.DeleteOne(context.Background(), filter)

		if err != nil {
			logging.FailedToDeleteOnDB(expense.ID, constants.CLOUD, err)
			return tagError.GetTagError(http.StatusBadRequest, err)
		}

		logging.DeletedOnDB(expense.ID, constants.CLOUD)
	}

	return nil
}

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

func ExpenseByCategory(ctx context.Context, cancel func(), errChan chan *tagError.TagError, userId string, month string, year int) (map[string][]domain.Expense, *tagError.TagError) {

	expVSCat := make(map[string][]domain.Expense, len(database.CacheDatabase.Categories))

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	// get the first and last day of the month
	startingDate, endingDate := timeHandler.GetFirstAndLastDayOfMonth(month, year)

	var wg sync.WaitGroup
	var mutex sync.Mutex

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

				mutex.Lock()
				defer mutex.Unlock()
				expVSCat[cat.Name] = expenses

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

func checkDuplicatedExpense(expense domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) bool {
	if statusDBLocal {
		var count int64
		database.DBlocal.Model(&domain.Expense{}).Where("description = ? AND value = ? AND paid_at = ? AND user_id = ?", expense.Description, expense.Value, expense.PaidAt, expense.UserID).Count(&count)

		if count > 0 {
			return true
		}
	}

	if statusDBCloud {
		filter := bson.M{"description": expense.Description, "value": expense.Value, "paidAt": expense.PaidAt, "userID": expense.UserID}

		count, err := database.DBCloud.Expense.CountDocuments(context.Background(), filter)

		if err != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Expense %s", expense.Description), constants.CLOUD, err)
			return true
		}

		if count > 0 {
			return true
		}
	}

	return false
}

func FindExpenseByDescription(description string, statusDBLocal bool, statusDBCloud bool) ([]domain.ExpenseDTO, *tagError.TagError) {
	if statusDBLocal {
		expenses := []domain.Expense{}

		result := database.DBlocal.Where("description = ?", description).Find(&expenses)

		if result.Error != nil {
			logging.FailedToFindOnDB(description, constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.FoundOnDB(description, constants.LOCAL)

		dto := []domain.ExpenseDTO{}

		for _, expense := range expenses {
			dto = append(dto, expense.ToDTO())
		}

		return dto, nil
	}

	if statusDBCloud {
		expenses := []domain.Expense{}

		objId, err := primitive.ObjectIDFromHex(description)

		if err != nil {
			logging.FailedToFindOnDB(description, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		filter := bson.M{"_id": objId}

		err = database.DBCloud.Expense.FindOne(context.Background(), filter).Decode(&expenses)

		if err != nil {
			logging.FailedToFindOnDB(description, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		logging.FoundOnDB(description, constants.CLOUD)

		dto := []domain.ExpenseDTO{}

		for _, expense := range expenses {
			dto = append(dto, expense.ToDTO())
		}

		return dto, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func FindExpenseByID(id string, statusDBLocal bool, statusDBCloud bool) (*domain.ExpenseDTO, *tagError.TagError) {
	if statusDBLocal {
		expense := domain.Expense{}

		result := database.DBlocal.Where("id = ?", id).First(&expense)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusBadRequest, result.Error)
		}

		logging.FoundOnDB(id, constants.LOCAL)

		dto := expense.ToDTO()

		return &dto, nil
	}

	if statusDBCloud {
		expense := domain.Expense{}

		objId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			logging.FailedToFindOnDB(id, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		filter := bson.M{"_id": objId}

		err = database.DBCloud.Expense.FindOne(context.Background(), filter).Decode(&expense)

		if err != nil {
			logging.FailedToFindOnDB(id, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		logging.FoundOnDB(id, constants.CLOUD)

		dto := expense.ToDTO()

		return &dto, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func makeExpenseParser(expense domain.ExpenseDTO) domain.ExpenseDTO {
	expenseParse := expense

	if expense.CategoryName != "" {
		category, tagErr := FindCategoryByName(expense.CategoryName, true, true)

		if tagErr != nil {
			logging.FailedToFindOnDB(fmt.Sprintf("Category %s", expense.CategoryName), "Category", tagErr.Inner)
			return expenseParse
		}

		if category.ID == "" {
			return expenseParse
		}

		expenseParse.CategoryID = category.ID
	}

	if expense.PaymentMethod != "" {
		cond := true

		for _, method := range constants.GetMethods() {
			if method == expense.PaymentMethod {
				cond = false
				break
			}
		}

		if cond {
			return expenseParse
		}

		expenseParse.PaymentMethod = expense.PaymentMethod
	}

	return expenseParse
}

func makeExpensesToUpdate(expenseParse domain.ExpenseDTO, statusDBLocal bool, statusDBCloud bool) ([]domain.ExpenseDTO, *tagError.TagError) {
	expensesToUpdate := []domain.ExpenseDTO{}

	expense, tagErr := FindExpenseByID(expenseParse.ID, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(expenseParse.ID, constants.LOCAL, tagErr.Inner)
		return nil, tagErr
	}

	if expense.TotalParcel > 0 {
		expenses, tagErr := FindExpenseByDescription(expense.Description, true, true)

		if tagErr != nil {
			logging.FailedToFindOnDB(expense.Description, constants.LOCAL, tagErr.Inner)
			return nil, tagErr
		}

		for _, exp := range expenses {
			if expenseParse.CategoryID != "" {
				exp.CategoryID = expenseParse.CategoryID
			}
			if expenseParse.PaymentMethod != "" {
				exp.PaymentMethod = expenseParse.PaymentMethod
			}
			if expenseParse.Value != 0 {
				exp.Value = expenseParse.Value
			}
			if expenseParse.Description != "" {
				exp.Description = expenseParse.Description
			}
			if expenseParse.PaidAt != "" {
				exp.PaidAt = expenseParse.PaidAt
			}

			expensesToUpdate = append(expensesToUpdate, exp)
		}
	} else {
		expensesToUpdate = append(expensesToUpdate, expenseParse)
	}

	return expensesToUpdate, nil
}
