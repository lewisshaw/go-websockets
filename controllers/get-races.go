package controllers

import (
	"lewisshaw/go-ws/models"

	"github.com/gin-gonic/gin"
)

type getRacesController struct {
	races models.RaceCollection
}

func NewGetRacesController(races models.RaceCollection) *getRacesController {
	return &getRacesController{
		races: races,
	}
}

func (grc *getRacesController) Action(c *gin.Context) {
	c.JSON(200, grc.races)
}
