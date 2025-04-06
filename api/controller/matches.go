package controller

import (
	"net/http"
	"strconv"

	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchesController struct {
	MatchesUseCase  		domain.MatchesUseCase
	Env 					*bootstrap.Env
}

func (mc *MatchesController) GetTeamMatches(c *gin.Context, collection *mongo.Collection){
	teamID := c.Param("id")
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	matches, err := mc.MatchesUseCase.GetTeamMatches(c, collection, limit, offset, teamID)

	if err != nil {
		c.JSON(http.StatusOK, domain.ErrMatchFetch)
		return
	}

	c.JSON(http.StatusOK, matches)
	return
}