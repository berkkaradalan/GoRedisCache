package usecase

import (
	"time"

	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/berkkaradalan/GoRedisCache/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMatchesUseCase(mr repository.MatchesRepository, timeout time.Duration, env *bootstrap.Env) *MatchesUseCase {
	return &MatchesUseCase{
		MatchesRepository: mr,
		contextTimeout:    timeout,
		Env:               env,
	}
}

type MatchesUseCase struct {
	MatchesRepository repository.MatchesRepository
	contextTimeout           time.Duration
	Env 			 *bootstrap.Env
}

func (mu *MatchesUseCase) GetTeamMatches(c *gin.Context, collection *mongo.Collection, limit int, offset int, teamID string) ([]domain.MatchResponse, error) {
	matches, err := mu.MatchesRepository.GetTeamMatches(c, collection, limit, offset, teamID)

	if err != nil && matches != nil {
		return matches, nil
	}

	var matchesResponses []domain.MatchResponse
	for _, match := range matches{
		matchesResponse := domain.MatchResponse{
			Area: match.Area,
			Competition: match.Competition,
			Season: match.Season,
			ID: match.ID,
			UTCDate: match.UTCDate,
			Status: match.Status,
			MatchDay: match.MatchDay,
			Stage: match.Stage,
			Group: match.Group,
			LastUpdated: match.LastUpdated,
			HomeTeam: match.HomeTeam,
			AwayTeam: match.AwayTeam,
			Score: match.Score,
			Refree: match.Refree,
		}
		matchesResponses = append(matchesResponses, matchesResponse)
	}

	return matchesResponses, nil
}