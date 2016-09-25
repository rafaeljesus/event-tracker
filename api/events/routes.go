package events

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/event-tracker/models"
	"net/http"
)

func Index(c echo.Context) error {
	return nil
}

func Create(c echo.Context) error {
	event := &models.Event{}

	if err := c.Bind(event); err != nil {
		return err
	}

	if err := event.Create(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &created{Ok: true})
}

type created struct {
	Ok bool `json:"ok"`
}
