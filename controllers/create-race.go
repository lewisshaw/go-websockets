package controllers

import (
	"fmt"
	"lewisshaw/go-ws/models"

	"github.com/gin-gonic/gin"
)

type createRaceController struct {
	races models.RaceCollection
}

func NewCreateRaceController(races models.RaceCollection) *createRaceController {
	return &createRaceController{
		races: races,
	}
}

func (grc *createRaceController) Action(c *gin.Context) {
	raceNum := grc.races.GetLength() + 1
	startSequence := models.NewStartSequence(60, []int{60, 30, 10, 0})
	race := models.NewRace(raceNum, fmt.Sprintf("Race %02d", raceNum), startSequence)

	grc.races.AddRace(race)

	c.JSON(200, race)
}
