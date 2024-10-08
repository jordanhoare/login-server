package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/tibia-oce/login-server/src/api"
	"github.com/tibia-oce/login-server/src/configs"
	grpc_login_server "github.com/tibia-oce/login-server/src/grpc"
	"github.com/tibia-oce/login-server/src/logger"
	"github.com/tibia-oce/login-server/src/server"
)

var numberOfServers = 2
var initDelay = 200

func main() {
	logger.Init(configs.GetLogLevel())
	logger.Info("Welcome to OTBR Login Server")

	var wg sync.WaitGroup
	wg.Add(numberOfServers)

	gConfigs := configs.GetGlobalConfigs()

	go startServer(&wg, gConfigs, grpc_login_server.Initialize(gConfigs))
	go startServer(&wg, gConfigs, api.Initialize(gConfigs))

	time.Sleep(time.Duration(initDelay) * time.Millisecond)
	gConfigs.Display()

	// wait until WaitGroup is done
	wg.Wait()
	logger.Info("Good bye...")
}

func startServer(
	wg *sync.WaitGroup,
	gConfigs configs.GlobalConfigs,
	server server.ServerInterface,
) {
	logger.Info(fmt.Sprintf("Starting %s server...", server.GetName()))
	logger.Error(server.Run(gConfigs))
	wg.Done()
	logger.Warn(fmt.Sprintf("Server %s is gone...", server.GetName()))
}
