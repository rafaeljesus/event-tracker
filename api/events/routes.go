package events

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/event-tracker/lib/kafka"
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

	msg, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := kafka.Enqueue("events", msg); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &response{Ok: true})
}

type response struct {
	Ok bool `json:"ok"`
}
