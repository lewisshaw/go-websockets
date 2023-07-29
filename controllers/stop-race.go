package controllers

import (
	"lewisshaw/go-ws/models"
	"lewisshaw/go-ws/sockets"
	"strconv"

	"github.com/gin-gonic/gin"
)

type stopRaceController struct {
	races        models.RaceCollection
	stopChannels sockets.StopRaceChannels
}

func NewStopRaceController(races models.RaceCollection, stopChannels sockets.StopRaceChannels) *stopRaceController {
	return &stopRaceController{
		races:        races,
		stopChannels: stopChannels,
	}
}

func (src *stopRaceController) Action(c *gin.Context) {
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

	src.stopChannels[id] <- true
}
