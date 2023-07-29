package main

import (
	"fmt"
	"lewisshaw/go-ws/controllers"
	"lewisshaw/go-ws/models"
	"lewisshaw/go-ws/sockets"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/gin-gonic/gin"
)

func main() {

	races := models.NewRaceCollection()
	clients := sockets.NewClients()
	stopRaceChannels := sockets.StopRaceChannels{}

	gin := gin.Default()
	gin.Use(corsMiddleware())

	gin.GET("/", controllers.Index)
	gin.GET("/races/view", controllers.NewGetRacesController(races).Action)
	gin.GET("/races/create", controllers.NewCreateRaceController(races).Action)
	gin.GET("/races/:id", controllers.NewViewRaceController(races).Action)
	gin.GET("/races/:id/ws", controllers.NewViewRaceSocketController(races, clients).Action)
	gin.GET("/races/:id/start", controllers.NewStartRaceController(clients, races, stopRaceChannels).Action)
	gin.GET("/races/:id/stop", controllers.NewStopRaceController(races, stopRaceChannels).Action)
	gin.POST("/races/:id/competitors", controllers.NewAddCompetitorToRaceController(races).Action)

	gin.Run()
}

func timer(w *widget.Label) {
	elapsed := 0
	for range time.Tick(time.Second) {
		elapsed++
		w.SetText(fmt.Sprint(elapsed))
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
