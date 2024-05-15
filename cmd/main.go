package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ruziba3vich/events/api/handlers"
	"github.com/ruziba3vich/events/postgres"
	"github.com/ruziba3vich/events/services"
)

func main() {
	router := gin.Default()

	db := postgres.DB()

	handler := handlers.New(&handlers.HandlerConfig{
		Services: services.NewService(db),
	})

	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", handler.LogInHandler)
	router.POST("/create-event", handler.CreateEventHandler)
	router.PUT("/update-event/:id", handler.UpdateEventHandler)
	router.DELETE("/delete-event/:id", handler.DeleteEventHandler)
	router.GET("/get-events", handler.GetAllEventsHandler)

	address := "localhost:7777"
	log.Println("Server is listening on", address)

	if err := router.Run(address); err != nil {
		log.Fatal("error starting the server")
	}
}
