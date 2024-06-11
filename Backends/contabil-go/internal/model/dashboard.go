package model

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/pkg/mathPlus"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
)

func GetDashboard(entry domain.DashboardPacket) (*domain.DashboardPacket, *tagError.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	ctx, cancel := context.WithCancel(context.Background())
	errors := make(chan *tagError.TagError)

	var wg sync.WaitGroup

	wg.Add(4)

	// TotalRanger routine for the income VS Expense plot and Resume data
	go totalRangeAndResumeRoutine(&wg, errors, ctx, cancel, &entry)

	// Timeline routine for the timeline
	go timelineRoutine(&wg, errors, ctx, cancel, &entry)

	// TODO: ExpenseByCategory
	go ExpenseByCategoryRoutine(&wg, errors, ctx, cancel, &entry)

	// TODO: ExpenseByMethod

	go ExpenseByMethodRoutine(&wg, errors, ctx, cancel, &entry)

	// TODO: FixatedExpenses

	wg.Wait()

	entry.UpdatedAt = timeHandler.GetTimeNow()

	return &entry, nil

}

func totalRangeAndResumeRoutine(wg *sync.WaitGroup, errChan chan *tagError.TagError, ctx context.Context, cancel func(), entry *domain.DashboardPacket) {

	defer wg.Done()
	ive, tagerr := TotalRanger(ctx, cancel, entry.UserID, entry.Month, entry.Year)

	if tagerr != nil {
		logging.GenericError("Failed to generate income vs expense", tagerr.Inner)
		errChan <- tagerr
		cancel()
	}

	logging.GenericSuccess(fmt.Sprintf("Incomes Vs Expenses from %s/%d for user %s sucessfully generated", entry.Month, entry.Year, entry.UserID))

	resume, tagerr := Resume(ive)

	if tagerr != nil {
		logging.GenericError("Failed to generate resume for dashboard", tagerr.Inner)
		errChan <- tagerr
		cancel()
	}

	logging.GenericSuccess(fmt.Sprintf("Resumes from %s/%d for user %s sucessfully generated", entry.Month, entry.Year, entry.UserID))
	entry.IncomevsExpense = ive
	entry.Resume = resume

}

func timelineRoutine(wg *sync.WaitGroup, errChan chan *tagError.TagError, ctx context.Context, cancel func(), entry *domain.DashboardPacket) {

	defer wg.Done()

	timeline, tagerr := Timeline(ctx, cancel, errChan, entry.UserID, entry.Month, entry.Year)

	if tagerr != nil {
		logging.GenericError("Failed to generate timeline array", tagerr.Inner)
		errChan <- tagerr
		cancel()
		return
	}

	entry.Timeline = timeline

}

func ExpenseByCategoryRoutine(wg *sync.WaitGroup, errChan chan *tagError.TagError, ctx context.Context, cancel func(), entry *domain.DashboardPacket) {

	defer wg.Done()

	ebc, tagErr := ExpenseByCategory(ctx, cancel, errChan, entry.UserID, entry.Month, entry.Year)

	if tagErr != nil {
		logging.GenericError("Failed to generate Expense by Category plot data", tagErr.Inner)
		errChan <- tagErr
		cancel()
		return
	}

	values := map[string]float64{}
	fixated := map[string][]domain.Activity{}

	fixe := []string{"Conta", "Streaming"}

	var wg2 sync.WaitGroup
	for catName, expenses := range ebc {

		wg2.Add(1)

		go expenseHandler(&wg2, ctx, catName, expenses, values, fixated, fixe)

	}

	wg2.Wait()

	entry.ExpenseByCategory = values
	entry.FixatedExpenses = fixated

}

func ExpenseByMethodRoutine(wg *sync.WaitGroup, errChan chan *tagError.TagError, ctx context.Context, cancel func(), entry *domain.DashboardPacket) {

	defer wg.Done()

	ebm, tagErr := ExpenseByMethod(ctx, cancel, errChan, entry.UserID, entry.Month, entry.Year)

	if tagErr != nil {
		logging.GenericError("Failed to generate Expense by Method plot data", tagErr.Inner)
		errChan <- tagErr
		cancel()
		return
	}

	entry.ExpenseByMethod = ebm

}

func expenseHandler(wg *sync.WaitGroup, ctx context.Context, catName string, expenses []domain.Expense, values map[string]float64, fixated map[string][]domain.Activity, fixe []string) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		logging.ContextAlreadyClosed()
		return

	default:
		value := 0.0

		for _, expense := range expenses {
			value += expense.Value

			for _, f := range fixe {
				if catName == f {
					fixated[catName] = append(fixated[catName], expense.ToActivity())
				}
			}
		}
		if value > 1 {
			values[catName] = mathPlus.ToFixed(value, 2)
		}
	}
}
