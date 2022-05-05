package app

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/middlewares"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"
	"codex/internal/pkg/utils/setter"

	"codex/internal/pkg/csrf"

	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunServer() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	// middleware
	middlewares.RegisterMetrics()
	metrics := middlewares.InitMetrics()
	api.Use(metrics.Metrics)

	api.Use(middlewares.Cors)
	api.Use(middlewares.Logger)
	api.Use(metrics.Metrics)
	api.Use(middlewares.PanicRecovery)
	// api.Use(middlewares.CsrfMdlw)

	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	setter.SetHandlers(setter.Services{
		Act: setter.Data{Db: db, Api: api},
		Mov: setter.Data{Db: db, Api: api},
		Usr: setter.Data{Db: db, Api: api},
		Col: setter.Data{Db: db, Api: api},
		Gen: setter.Data{Db: db, Api: api},
		Ann: setter.Data{Db: db, Api: api},
		Ser: setter.Data{Db: db, Api: api},
		Pla: setter.Data{Db: db, Api: api},

		Com: setter.Data{Db: nil, Api: api},
		Rat: setter.Data{Db: nil, Api: api},
		Aut: setter.Data{Db: nil, Api: api},
	})
	router.Handle("/metrics", promhttp.Handler())

	csrfsecurity.SetCsrf(api)

	port := os.Getenv("PORT") // to get port from Heroku
	if port == "" {
		port = config.DevConfigStore.LocalPort
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	log.Info("Connecting to port " + port)

	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
	}
}
