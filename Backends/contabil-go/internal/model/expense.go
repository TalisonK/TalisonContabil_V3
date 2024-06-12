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

func UpdateExpenseHandler(expense domain.ExpenseDTO) ([]string, *tagError.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if expense.ID == "" {
		return nil, tagError.GetTagError(http.StatusBadRequest, logging.InvalidFields())
	}

	expenseParse, tagErr := FindExpenseByDescription(expense.Description, statusDBLocal, statusDBCloud)

	if tagErr != nil {
		logging.FailedToFindOnDB(expense.ID, constants.LOCAL, tagErr.Inner)
		return nil, tagErr
	}

	fmt.Print(expenseParse)

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())

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

		result := database.DBlocal.Where("de = ?", description).Find(&expenses)

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
