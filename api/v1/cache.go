package v1

import "github.com/labstack/echo/v4"

type CacheControllerInf interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Get(c echo.Context) error
}

type CacheControllerInstance struct {
}

func CacheController() CacheControllerInf {
	return new(CacheControllerInstance)
}

func (e CacheControllerInstance) Create(c echo.Context) error {
	panic("implement me")
}

func (e CacheControllerInstance) Update(c echo.Context) error {
	panic("implement me")
}

func (e CacheControllerInstance) Delete(c echo.Context) error {
	panic("implement me")
}

func (e CacheControllerInstance) Get(c echo.Context) error {
	panic("implement me")
}
