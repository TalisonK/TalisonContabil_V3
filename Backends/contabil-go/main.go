package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TalisonK/TalisonContabil/src/auth"
	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/routes"
	"github.com/TalisonK/TalisonContabil/src/util/constants"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
)

func main() {

	err := config.Load()

	auth.NewAuth()

	if err != nil {
		log.Fatal("Erro ao carregar as configurações, certifique-se de que o arquivo de configuração está correto.")
		logging.GenericError("Erro ao carregar as configurações", err)
		os.Exit(2)
	}

	err = database.OpenConnectionLocal()

	if err != nil {
		logging.FailedToOpenConnection(constants.LOCAL, err)
	}

	err = database.OpenConnectionCloud()

	if err != nil {
		logging.FailedToOpenConnection(constants.CLOUD, err)
	}

	defer database.CloseConnections()

	// Create a new Chi router

	r := routes.Router()

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
