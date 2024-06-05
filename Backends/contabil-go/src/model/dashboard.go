package model

import (
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func GetDashboard(entry domain.DashboardPacket) (*domain.DashboardPacket, *util.TagError) {

	//errors := make(chan[] *util.TagError)

	// initializing package variables
	entry.IncomevsExpense = make([]domain.IncomevsExpense, 13)

	//var ws sync.WaitGroup

	ive, tagerr := TotalRanger(entry.UserID, entry.Month, entry.Year)

	if tagerr != nil {
		logging.GenericError("Failed to generate income vs expense", tagerr.Inner)
		return nil, tagerr
	}

	entry.IncomevsExpense = ive

	timeline, tagerr := Timeline(entry.UserID, entry.Month, entry.Year)

	if tagerr != nil {
		logging.GenericError("Failed to generate timeline array", tagerr.Inner)
		return nil, tagerr
	}

	entry.Timeline = timeline

	entry.UpdatedAt = util.GetTimeNow()

	return &entry, nil

}
