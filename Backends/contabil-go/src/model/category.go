package model

import (
	"fmt"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
)

func GetCategories() ([]domain.Category, error) {

	statusDBLocal, statusDBCloud := util.CheckDBStatus()

	if !statusDBLocal && !statusDBCloud {
		util.LogHandler("No database connection available.", nil, "model.CreateUser")
		return nil, fmt.Errorf("no database connection available")
	}

	if statusDBLocal {

		var categories []domain.Category
		if result := database.DBlocal.Find(&categories); result.Error != nil {
			util.LogHandler("Failed to get categories from local database", result.Error, "model.GetCategories")
			return nil, result.Error
		} else {
			util.LogHandler("Categories retrieved from local database", nil, "model.GetCategories")
			return categories, nil
		}
	}
	return nil, nil
}
