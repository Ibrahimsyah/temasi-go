package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewCommonRouter(e *echo.Echo) *echo.Group {
	router := e.Group("/")

	router.GET("ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})

	return router
}
