package mcsauth

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/utils/log"
	proto "codex/internal/pkg/authorization/delivery/grpc"
	"codex/internal/pkg/authorization/repository"
	"codex/internal/pkg/authorization/usecase"

	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	autRep := autrepository.InitAutRep(db)
	autUsc := autusecase.InitAutUsc(autRep)

	s := grpc.NewServer()

	proto.RegisterAutherServer(s, autUsc)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Warn("{RunServer} mcs auth")
		log.Error(err)
	}

	s.Serve(l)
}
