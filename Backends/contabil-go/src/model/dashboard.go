package model

import (
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func GetDashboard(entry domain.DashboardPacket) (*domain.DashboardPacket, *util.TagError) {

	ive, tagerr := TotalRanger(entry.UserID, entry.Month, entry.Year)

	if tagerr != nil {
		logging.GenericError("Failed to generate income vs expense", tagerr.Inner)
		return nil, tagerr
	}

	entry.IncomevsExpense = ive

	entry.UpdatedAt = util.GetTimeNow()

	return &entry, nil

}
