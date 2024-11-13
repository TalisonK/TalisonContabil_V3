package model

import (
	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
)

func StartCache() database.CacheDB {

	cache := database.CacheDB{}

	cats, tagerr := GetCategories()

	if tagerr != nil {
		logging.FailedToFindOnDB("All Categories", constants.LOCAL, tagerr.Inner)
		cache.CategoryStatus = false
	} else {
		cache.CategoryStatus = true
		cache.Categories = map[string]domain.Category{}

		for _, category := range cats {
			cache.Categories[category.ID] = category
		}
	}

	logging.GenericSuccess("Cache for Category started successfully.")

	cards, tagerr := GetCreditCards()

	if tagerr != nil {
		logging.FailedToFindOnDB("All Credit Cards", constants.LOCAL, tagerr.Inner)
		cache.CardStatus = false
	} else {
		cache.CardStatus = true
		cache.Cards = map[string]domain.CreditCard{}

		for _, card := range cards {
			cache.Cards[card.ID] = card
		}
	}

	logging.GenericSuccess("Cache for Credit Card started successfully.")

	return cache

}
