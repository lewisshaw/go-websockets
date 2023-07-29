package controllers

import (
	"fmt"
	"lewisshaw/go-ws/models"
	"lewisshaw/go-ws/sockets"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type viewRaceSocketController struct {
	races    models.RaceCollection
	upgrader *websocket.Upgrader
	clients  sockets.Clients
}

func NewViewRaceSocketController(races models.RaceCollection, clients sockets.Clients) *viewRaceSocketController {
	return &viewRaceSocketController{
		races: races,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: clients,
	}
}

func (vrc *viewRaceSocketController) Action(c *gin.Context) {
	conn, err := vrc.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		return
	}
	race := vrc.races.GetRaceById(id)

	if nil == race {
		log.Println(fmt.Sprintf("Could not find race for ID %d", id))
		return
	}

	vrc.clients.AddClient(sockets.NewClient(conn, race))
}
