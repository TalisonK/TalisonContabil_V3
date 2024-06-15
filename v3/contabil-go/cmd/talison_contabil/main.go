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

	// Remove the log file
	os.Remove("talisoncontabil.log")

	// Load the configuration file
	err := config.Load()

	if err != nil {
		log.Fatal("Erro ao carregar as configurações, certifique-se de que o arquivo de configuração está correto.")
		logging.GenericError("Erro ao carregar as configurações", err)
		os.Exit(2)
	}

	// Start the trace
	if !config.IsProd() {
		f, err := os.Create("trace.out")

		if err != nil {
			logging.GenericError("Could no create tracer file", err)
		}

		err = trace.Start(f)

		if err != nil {
			logging.GenericError("Could no Start tracer", err)
		}

		defer trace.Stop()
	}

	// Start the auth
	auth.NewAuth()

	// Open the database Local connections
	err = database.OpenConnectionLocal()

	if err != nil {
		logging.FailedToOpenConnection(constants.LOCAL, err)
	}

	// Open the database Cloud connections
	err = database.OpenConnectionCloud()

	if err != nil {
		logging.FailedToOpenConnection(constants.CLOUD, err)
	}

	// Start the cache
	database.CacheDatabase = model.StartCache()

	// Close the database connections
	defer database.CloseConnections()

	// Create a new Chi router
	r := routes.Router()

	// Start the server
	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
