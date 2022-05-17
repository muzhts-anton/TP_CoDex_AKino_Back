package mcscomt

import (
	proto "codex/internal/pkg/comment/delivery/grpc"
	comrepository "codex/internal/pkg/comment/repository"
	comusecase "codex/internal/pkg/comment/usecase"
	"codex/internal/pkg/database"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"

	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	autRep := comrepository.InitComRep(db)
	autUsc := comusecase.InitComUsc(autRep)

	s := grpc.NewServer()

	proto.RegisterPosterServer(s, autUsc)

	l, err := net.Listen(config.DevConfigStore.Mcs.Comment.ConnType, ":"+config.DevConfigStore.Mcs.Comment.Port)
	if err != nil {
		log.Warn("{RunServer} mcs comt")
		log.Error(err)
	}

	s.Serve(l)
}
