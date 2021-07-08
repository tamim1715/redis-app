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
		log.Fatal("database error: " + err.Error())
	}
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
	defer log.Fatal("Server died")
}
