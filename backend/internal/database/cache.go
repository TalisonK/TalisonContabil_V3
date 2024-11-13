package database

import "github.com/TalisonK/TalisonContabil/internal/domain"

type CacheDB struct {
	Categories     map[string]domain.Category
	Cards          map[string]domain.CreditCard
	CategoryStatus bool
	CardStatus     bool
}

var CacheDatabase CacheDB

func (c CacheDB) GetCategoryById(id string) domain.Category {

	return c.Categories[id]

}

func (c CacheDB) GetCardById(id string) domain.CreditCard {

	return c.Cards[id]

}
