package app

import (
	"codex/internal/pkg/middlewares"
	"codex/internal/pkg/utils/log"
	"codex/internal/pkg/database"

	"codex/internal/pkg/user/repository"
	"codex/internal/pkg/user/usecase"
	"codex/internal/pkg/user/delivery"

	"codex/internal/pkg/collections/repository"
	"codex/internal/pkg/collections/usecase"
	"codex/internal/pkg/collections/delivery"

	"codex/internal/pkg/movie/repository"
	"codex/internal/pkg/movie/usecase"
	"codex/internal/pkg/movie/delivery"

	"codex/internal/pkg/actor/repository"
	"codex/internal/pkg/actor/usecase"
	"codex/internal/pkg/actor/delivery"

	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func RunServer() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(middlewares.Cors)

	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	actRep := actrepository.InitActRep(db)
	movRep := movrepository.InitMovRep(db)
	usrRep := usrrepository.InitUsrRep(db)
	colRep := colrepository.InitColRep(db)

	actUsc := actusecase.InitActUsc(actRep)
	movUsc := movusecase.InitMovUsc(movRep)
	usrUsc := usrusecase.InitUsrUsc(usrRep)
	colUsc := colusecase.InitColUsc(colRep)

	actdelivery.SetActHandlers(api, actUsc)
	movdelivery.SetMovHandlers(api, movUsc)
	usrdelivery.SetUsrHandlers(api, usrUsc)
	coldelivery.SetColHandlers(api, colUsc)

	port := os.Getenv("PORT") // to get port from Heroku
	if port == "" {
		port = "3000"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	log.Info("connecting to port " + port)

	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
	}
}
