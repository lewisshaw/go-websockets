package controllers

import (
	"lewisshaw/go-ws/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type viewRaceController struct {
	races models.RaceCollection
}

func NewViewRaceController(races models.RaceCollection) *viewRaceController {
	return &viewRaceController{
		races: races,
	}
}

func (vrc *viewRaceController) Action(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be provided and must be an int",
		})
		return
	}
	race := vrc.races.GetRaceById(id)

	if nil == race {
		c.JSON(404, gin.H{
			"error": "race not found for id",
		})
		return
	}

	c.JSON(200, race)
}
