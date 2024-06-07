package model

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func GetDashboard(entry domain.DashboardPacket) (*domain.DashboardPacket, *util.TagError) {

	statusDBLocal, statusDBCloud := database.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	ctx, cancel := context.WithCancel(context.Background())
	errors := make(chan *util.TagError)

	var wg sync.WaitGroup

	wg.Add(2)

	// TotalRanger routine for the income VS Expense plot and Resume data
	go totalRangeAndResumeRoutine(&wg, errors, ctx, cancel, &entry)

	// Timeline routine for the timeline
	go timelineRoutine(&wg, errors, ctx, cancel, &entry)

	// TODO: ExpenseByCategory

	// TODO: ExpenseByMethod

	// TODO: FixatedExpenses

	wg.Wait()

	entry.UpdatedAt = util.GetTimeNow()

	return &entry, nil

}

func totalRangeAndResumeRoutine(wg *sync.WaitGroup, errChan chan *util.TagError, ctx context.Context, cancel func(), entry *domain.DashboardPacket) {

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

func timelineRoutine(wg *sync.WaitGroup, errChan chan *util.TagError, ctx context.Context, cancel func(), entry *domain.DashboardPacket) {

	defer wg.Done()

	timeline, tagerr := Timeline(ctx, cancel, errChan, entry.UserID, entry.Month, entry.Year)

	if tagerr != nil {
		logging.GenericError("Failed to generate timeline array", tagerr.Inner)
		errChan <- tagerr
		cancel()
	}

	entry.Timeline = timeline

}
