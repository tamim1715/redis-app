package main

import (
	"github.com/khan1507017/redis-app/config"
	"github.com/khan1507017/redis-app/database/rds"
	"github.com/khan1507017/redis-app/router"
	"github.com/khan1507017/redis-app/server"
	"github.com/prometheus/common/log"
)

func main() {

	srv := server.New()
	router.Routes(srv)
	err := config.InitEnvironmentVariables()
	if err != nil {
		log.Fatal("envVars error: " + err.Error())
	}
	err = rds.InitRedisMaster()
	if err != nil {
		log.Fatal("master endpoint error: " + err.Error())
	}
	err = rds.InitRedisSlave()
	if err != nil {
		log.Fatal("slave endpoint error: " + err.Error())
	}
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
