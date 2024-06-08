package model

import (
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func StartCache() database.CacheDB {

	cache := database.CacheDB{}

	cat, tagerr := GetCategories()

	if tagerr != nil {
		logging.FailedToFindOnDB("All Categories", constants.LOCAL, tagerr.Inner)
		cache.CategoryStatus = false
	} else {
		cache.Categories = cat
		cache.CategoryStatus = true
	}

	return cache

}
