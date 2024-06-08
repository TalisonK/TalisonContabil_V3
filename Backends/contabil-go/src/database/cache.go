package database

import "github.com/TalisonK/TalisonContabil/src/domain"

type CacheDB struct {
	Categories     []domain.Category
	CategoryStatus bool
}

var CacheDatabase CacheDB
