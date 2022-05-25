package app

import (
	announcedRepository "codex/internal/pkg/announced/repository"
	csrfsecurity "codex/internal/pkg/csrf"
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/middlewares"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"
	"codex/internal/pkg/utils/setter"
	"strconv"

	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

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
		// announcedRepo: setter.Data{Db: db, Api: api},

		Com: setter.Data{Db: nil, Api: api},
		Rat: setter.Data{Db: nil, Api: api},
		Aut: setter.Data{Db: nil, Api: api},
	})
	router.Handle("/metrics", promhttp.Handler())

	announcedRepo := announcedRepository.InitAnnRep(db)
	go notificationWorker(announcedRepo)

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

func notificationWorker(announcedRepo domain.AnnouncedRepository) {
	// storage for announceds released today
	comingAnnounced := struct {
		announceds []domain.Announced
		sync.RWMutex
	}{}

	// daemon to update coming this month announceds
	go func(announcedRepo domain.AnnouncedRepository) {
		for {
			comingAnnounced.Lock()
			comingAnnounced.announceds = []domain.Announced{}
			comingAnnounced.Unlock()

			year, month, _ := time.Now().Date()
			announcedsBuffer, err := announcedRepo.GetAnnouncedByMonthYear(int(month), year)
			if err == nil {
				for _, v := range announcedsBuffer.AnnouncedList {
					log.Info(fmt.Sprintf("AnnouncedList = %d ", len(announcedsBuffer.AnnouncedList)))
					log.Info("Releasedate" + string(v.Releasedate))
					log.Info("time.Now()" + time.Now().Format("2006-01-02"))
					if time.Now().Format("2006-01-02") == v.Releasedate {
						comingAnnounced.Lock()
						comingAnnounced.announceds = append(comingAnnounced.announceds, v)
						comingAnnounced.Unlock()
					}
				}
			}
			log.Info(fmt.Sprintf("Found %d announceds released today", len(comingAnnounced.announceds)))
			time.Sleep(10 * time.Second) // TO DO 24 hours 
		}
	}(announcedRepo)

	// preparing firebase messages sender
	opt := option.WithCredentialsFile("firebasePrivateKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Warn(fmt.Sprintf("error initializing Firebase app: %v\n", err))
	}
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Warn(fmt.Sprintf("error getting Firebase Messaging client: %v\n", err))
	}

	// then we send notifications for all announceds that were released today
	ticker := time.NewTicker(time.Minute)
	for {
		select{
		case <-ticker.C:
			// hr, min, _ := time.Now().Clock()
			if true {
			// if hr == 12 && min == 0 {
				comingAnnounced.RLock()
				for _, v := range comingAnnounced.announceds {
					message := &messaging.Message{
						Notification: &messaging.Notification{
							Title: "Сегодня премьера фильма",
							Body:  v.Title,
						},
						Topic: "all",
					}
					log.Info(message.Topic)
					response, err := client.Send(ctx, message)
					if err != nil {
						log.Error(err)
					}
					id, err := strconv.Atoi(v.Id)
					if err != nil{
						log.Error(err)
						return
					}
					log.Info(fmt.Sprintf("Successfully sent message: %v, for announced id: %d", response, id))
				}
				comingAnnounced.RUnlock()
			}
			time.Sleep(24 * time.Hour)
		}
	}
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		hr, min, _ := time.Now().Clock()
	// 		if hr == 12 && min == 0 {
	// 			comingFilms.RLock()
	// 			for _, v := range comingFilms.films {
	// 				message := &messaging.Message{
	// 					Notification: &messaging.Notification{
	// 						Title: "Сегодня вышел в прокат фильм",
	// 						Body:  v.Title,
	// 					},
	// 					Topic: "all",
	// 				}
	// 				response, err := client.Send(ctx, message)
	// 				if err != nil {
	// 					log.Error(err)
	// 				}
	// 				log.Info(fmt.Sprintf("Successfully sent message: %v, for film id: %d", response, v.Id))
	// 			}
	// 			comingFilms.RUnlock()
	// 		}
	// 	}
	// }
}
