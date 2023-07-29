package controllers

import (
	"lewisshaw/go-ws/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addCompetitorToRaceController struct {
	races models.RaceCollection
}

func NewAddCompetitorToRaceController(races models.RaceCollection) *addCompetitorToRaceController {
	return &addCompetitorToRaceController{
		races: races,
	}
}

type competitorRequest struct {
	Name string `json:"name" binding:"required"`
}

func (acr *addCompetitorToRaceController) Action(c *gin.Context) {
	race, ok := getRaceFromUrlIdParam(c.Param("id"), c, acr.races)
	if !ok {
		return
	}
	var json competitorRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	competitor := models.NewCompetitor(race.GetCompetitorCount(), json.Name)
	race.AddCompetitors(competitor)
	c.JSON(200, competitor)
}
