package grpc_login_server

import (
	"database/sql"
	"net"

	"github.com/tibia-oce/login-server/src/configs"
	"github.com/tibia-oce/login-server/src/database"
	"github.com/tibia-oce/login-server/src/grpc/login_proto_messages"
	"github.com/tibia-oce/login-server/src/server"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	DB *sql.DB
	login_proto_messages.LoginServiceServer
	server.ServerInterface
}

func Initialize(gConfigs configs.GlobalConfigs) *GrpcServer {
	var ls GrpcServer

	ls.DB = database.PullConnection(gConfigs)

	return &ls
}

func (ls *GrpcServer) Run(gConfigs configs.GlobalConfigs) error {
	c, err := net.Listen("tcp", gConfigs.LoginServerConfigs.Grpc.Format())

	if err != nil {
		return err
	}

	server := grpc.NewServer()
	login_proto_messages.RegisterLoginServiceServer(server, ls)

	return server.Serve(c)
}

func (ls *GrpcServer) GetName() string {
	return "gRPC"
}
