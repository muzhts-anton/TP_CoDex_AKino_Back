package mcsauth

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/utils/log"
	_ "codex/internal/pkg/authorization/repository"
	
	proto "codex/proto"

	"net"
	_ "context"

	"google.golang.org/grpc"
)

type GRPCServer struct{
	proto.UnimplementedGetUserServer
}

func RunServer() {
	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	s := grpc.NewServer()
	srv := &GRPCServer{}

	proto.RegisterGetUserServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Warn("{RunServer} mcs auth")
		log.Error(err)
	}

	s.Serve(l)
}
