package healthz

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{Alive: true})
}

type response struct {
	Alive bool `json:"alive"`
}
