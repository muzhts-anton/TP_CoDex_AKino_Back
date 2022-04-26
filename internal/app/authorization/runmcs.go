package mcsauth

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/utils/log"
	"codex/internal/pkg/authorization/repository"
	
	proto "codex/proto"

	"net"
	"context"

	"google.golang.org/grpc"
)

type GRPCServer struct{
	proto.UnimplementedGetUserServer
}

func (s *GRPCServer) GetById(ctx context.Context, req *proto.Id) (*proto.User, error) {
	autrepository.
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
