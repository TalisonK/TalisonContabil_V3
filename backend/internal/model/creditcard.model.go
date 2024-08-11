package model

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
)

func CreateCreditCard() (*domain.CreditCard, *tagError.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	return nil, nil

}
