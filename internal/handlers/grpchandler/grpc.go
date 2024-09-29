package grpchandler

import (
	"fmt"
	"net"

	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/core/service/app"
	"github.com/TomasCruz/users/internal/handlers/grpchandler/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCHandler struct {
	svc    app.AppUserService
	logger ports.Logger
	server *grpc.Server
	users.UnimplementedUsersServer
}

func New(port string, svc app.AppUserService, logger ports.Logger) *GRPCHandler {
	// create a listener on TCP port
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		logger.Fatal(err, "gRPC listener error")
	}

	// create and register gRPC server
	grpcHandler := GRPCHandler{svc: svc, logger: logger, server: grpc.NewServer()}
	users.RegisterUsersServer(grpcHandler.server, &grpcHandler)
	reflection.Register(grpcHandler.server)

	// fire up the gRPC server
	go func() {
		logger.Info(nil, fmt.Sprintf("starting gRPC server on :%s", port))
		err := grpcHandler.server.Serve(listener)
		if err != nil {
			logger.Fatal(err, "gRPC Serve error")
		}
	}()

	return &grpcHandler
}
