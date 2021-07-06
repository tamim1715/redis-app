package main

import (
	"github.com/khan1507017/redis-app/config"
	"github.com/khan1507017/redis-app/router"
	"github.com/khan1507017/redis-app/server"
)

func main() {

	srv := server.New()
	router.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
