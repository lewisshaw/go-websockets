package controllers

import (
	"fmt"
	"lewisshaw/go-ws/models"
	"lewisshaw/go-ws/sockets"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type startRaceController struct {
	clients      sockets.Clients
	races        models.RaceCollection
	stopChannels sockets.StopRaceChannels
}

func NewStartRaceController(clients sockets.Clients, races models.RaceCollection, stopChannels sockets.StopRaceChannels) *startRaceController {
	return &startRaceController{
		clients:      clients,
		races:        races,
		stopChannels: stopChannels,
	}
}

func (src *startRaceController) Action(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be provided and must be an int",
		})
		return
	}

	race := src.races.GetRaceById(id)

	if nil == race {
		c.JSON(404, gin.H{
			"error": "race not found for id",
		})
		return
	}
	go src.startRace(id, race)
}

func (src *startRaceController) startRace(id int, race models.Race) {
	hootSeconds := 3
	hooting := false
	race.GetStartSequence().Start()
outer:
	for {
		select {
		case remainingTime := <-race.GetStartSequence().TickChannel():
			if hooting && hootSeconds == 0 {
				hooting = false
				hootSeconds = 3
			}
			if hooting {
				hootSeconds--
			}
			src.clients.MessageClientsForRaceIdStarting(race.GetId(), remainingTime, hooting)
		case remainingTime := <-race.GetStartSequence().HootChannel():
			hooting = true
			src.clients.MessageClientsForRaceIdStarting(race.GetId(), remainingTime, hooting)
			fmt.Println("Hooting")
		case <-race.GetStartSequence().CompleteChannel():
			break outer
		}

	}
	timer := time.NewTicker(time.Second)
	stopChannel := make(chan bool)
	src.stopChannels[id] = stopChannel
	go handleTimer(race, timer, src.clients, stopChannel)
}

func handleTimer(race models.Race, timer *time.Ticker, clients sockets.Clients, stopChannel chan bool) {
	for {
		select {
		case <-timer.C:
			race.IncrementElapsedTime()
			clients.MessageClientsForRaceId(race.GetId())
		case <-stopChannel:
			timer.Stop()
			return
		}
	}
}
