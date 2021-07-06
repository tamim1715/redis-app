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
	err := rds.InitRedisMaster()
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
