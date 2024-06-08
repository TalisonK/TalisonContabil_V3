package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"

	"github.com/TalisonK/TalisonContabil/internal/auth"
	"github.com/TalisonK/TalisonContabil/internal/config"
	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/internal/model"
	"github.com/TalisonK/TalisonContabil/internal/routes"
)

func main() {

	f, err := os.Create("trace.out")

	if err != nil {
		logging.GenericError("Could no create tracer file", err)
	}

	err = trace.Start(f)

	if err != nil {
		logging.GenericError("Could no Start tracer", err)
	}

	defer trace.Stop()

	err = config.Load()

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

	database.CacheDatabase = model.StartCache()

	defer database.CloseConnections()

	// Create a new Chi router

	r := routes.Router()

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
