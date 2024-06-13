package database

import "github.com/TalisonK/TalisonContabil/internal/domain"

type CacheDB struct {
	Categories     map[string]domain.Category
	CategoryStatus bool
}

var CacheDatabase CacheDB

func (c CacheDB) GetNameById(id string) domain.Category {

	return c.Categories[id]

}
