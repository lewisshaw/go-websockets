package controllers

import (
	"lewisshaw/go-ws/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getRaceFromUrlIdParam(idString string, c *gin.Context, races models.RaceCollection) (models.Race, bool) {

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id must be provided and must be an int",
		})
		return nil, false
	}

	race := races.GetRaceById(id)

	if nil == race {
		c.JSON(404, gin.H{
			"error": "race not found for id",
		})
		return nil, false
	}

	return race, true
}
