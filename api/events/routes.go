package events

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rafaeljesus/event-tracker/lib/kafka"
	"github.com/rafaeljesus/event-tracker/models"
	"net/http"
)

func Index(c echo.Context) error {
	name := c.QueryParam("name")
	query := models.Query{name}

	err, result := models.Search(query)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func Create(c echo.Context) error {
	event := &models.Event{}
	if err := c.Bind(event); err != nil {
		return err
	}

	payload, _ := json.Marshal(event)
	if err := kafka.Enqueue("events", payload); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, &response{Ok: true})
}

type response struct {
	Ok bool `json:"ok"`
}
