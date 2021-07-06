package router

import (
	v1 "github.com/khan1507017/redis-app/api/v1"
	"github.com/labstack/echo/v4"
)

func V1_routes(group *echo.Group) {
	group.POST("/cache", v1.CacheController().Create)
	group.GET("/cache", v1.CacheController().Get)
	group.PUT("/cache", v1.CacheController().Update)
	group.DELETE("/cache", v1.CacheController().Delete)
}
