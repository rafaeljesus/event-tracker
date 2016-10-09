package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/event-tracker/api/events"
	"github.com/rafaeljesus/event-tracker/api/healthz"
	"github.com/rafaeljesus/event-tracker/lib/elastic"
	"github.com/rafaeljesus/event-tracker/lib/kafka"
	"github.com/rafaeljesus/event-tracker/listener"
	"log"
	"os"
)

func main() {
	log.Print("Connecting Elastic...")
	elastic.Connect()

	log.Print("Connecting Kafka...")
	kafka.Connect()
	kafka.FromQueue("events", listener.EventCreated)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	r := e.Group("/v1")

	r.GET("/healthz", healthz.Index)
	r.GET("/events", events.Index)
	r.POST("/events", events.Create)

	log.Print("Server starting...")
	e.Run(fasthttp.New(":" + os.Getenv("EVENT_TRACKER_PORT")))
}
