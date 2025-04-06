package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/berkkaradalan/GoRedisCache/repository"
	"github.com/berkkaradalan/GoRedisCache/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

 func NewTeamsUseCase(tr repository.TeamsRepository, timeout time.Duration, env *bootstrap.Env) *TeamsUseCase {
	return &TeamsUseCase{
		teamsRepository: 			tr,
		contextTimeout: 			timeout,
		Env: 						env,
 	}
}

type TeamsUseCase struct{
 	teamsRepository 			repository.TeamsRepository
 	contextTimeout 				time.Duration
 	Env 						*bootstrap.Env
}

func (tu *TeamsUseCase) GetTeamById(ctx *gin.Context, collection *mongo.Collection, redisClient *redis.Client,team_id string) (domain.TeamResponse, error) {
	cacheKey := "team:" + team_id
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedTeam domain.TeamResponse
		if jsonErr := json.Unmarshal([]byte(cachedData), &cachedTeam); jsonErr == nil {
			return cachedTeam, nil
		}
	}
	
	team, err := tu.teamsRepository.GetTeamById(ctx, collection, team_id)

	if err == nil && team != nil {
		return *team, nil
	}

	resp, err := utils.RequestToExternalApi(tu.Env.FOOTBALL_DATA_API_KEY, tu.Env.FOOTBALL_DATA_TEAM_API_URL+team_id)
	if err != nil {
		return domain.TeamResponse{}, fmt.Errorf("no team found with id: %s", team_id)
	}

	if errMsg, exists := resp["error"]; exists {
		if errCode, ok := errMsg.(float64); ok && int(errCode) == 404 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Team not found in external API"})
			return domain.TeamResponse{}, fmt.Errorf("team with id %s not found", team_id)
		}
	}

	if errMsg, exists := resp["errorCode"]; exists {
		if errCode, ok := errMsg.(float64); ok && int(errCode) == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Team Id is not valid"})
			return domain.TeamResponse{}, fmt.Errorf("team with id %s not valid", team_id)
		}
	}

	var teamArea *domain.TeamArea
	if area, ok := resp["area"].(map[string]interface{}); ok {
		teamArea = &domain.TeamArea{
			ID:       utils.GetInt(bson.M(area), "id"),
			Name:     utils.GetString(bson.M(area), "name"),
			CountryCode: utils.GetString(bson.M(area), "code"),
			FlagUrl: utils.GetStringPointer(bson.M(area), "flag"),
		}
	}

	var squad []*domain.TeamSquad
	if squadList, ok := resp["squad"].([]interface{}); ok {
		for _, item := range squadList {
			if squadMap, valid := item.(map[string]interface{}); valid {
				squad = append(squad, &domain.TeamSquad{
					ID:        	 utils.GetInt(bson.M(squadMap), "id"),
					Name:        utils.GetString(bson.M(squadMap), "name"),
					Position:    utils.GetString(bson.M(squadMap), "position"),
					DateOfBirth: utils.GetString(bson.M(squadMap), "dateOfBirth"),
					Nationality: utils.GetString(bson.M(squadMap), "nationality"),
				})
			}
		}
	}

	var competitions []*domain.TeamCompetition
	if runningCompetitions, ok := resp["runningCompetitions"].([]interface{}); ok {
		for _, compItem := range runningCompetitions {
			if compMap, valid := compItem.(map[string]interface{}); valid {
				competitions = append(competitions, &domain.TeamCompetition{
					ID:       utils.GetInt(bson.M(compMap), "id"),
					Name:     utils.GetString(bson.M(compMap), "name"),
					Code:     utils.GetString(bson.M(compMap), "code"),
					Type:     utils.GetString(bson.M(compMap), "type"),
					EmblemUrl: utils.GetStringPointer(bson.M(compMap), "emblem"),
				})
			}
		}
	}

	var lastUpdated *string
	if val, ok := resp["lastUpdated"].(string); ok {
		lastUpdated = &val
	}

	convertedTeam := domain.TeamResponse{
		Area: teamArea,
		ID: utils.GetInt(resp, "id"),
		Name: utils.GetString(resp, "name"),
		ShortName: utils.GetString(resp, "shortName"),
		Tla: utils.GetString(resp, "tla"),
		CrestUrl: utils.GetStringPointer(resp, "crestUrl"),
		Address: utils.GetString(resp, "address"),
		Website: utils.GetString(resp, "website"),
		Founded: utils.GetInt(resp, "founded"),
		ClubColors: utils.GetString(resp, "clubColors"),
		Venue: utils.GetString(resp, "venue"),
		Competition: competitions,
		Squad: squad,
		LastUpdated: lastUpdated,
	}

	_, err = collection.InsertOne(ctx, convertedTeam)
	if err != nil {
		log.Printf("MongoDB insert error: %v", err)
	}

	cacheTeam(ctx, redisClient, cacheKey, &convertedTeam)
	return convertedTeam, nil
}

func cacheTeam(ctx context.Context, redisClient *redis.Client, key string, team *domain.TeamResponse) {
	teamJSON, err := json.Marshal(team)
	if err != nil {
		log.Printf("Failed to marshal team data: %v", err)
		return
	}
	status := redisClient.Set(ctx, key, teamJSON, time.Hour)
	if status.Err() != nil {
        log.Printf("Redis set error: %v", status.Err())
    } else {
        log.Printf("Successfully cached team: %s", key)
    }
}

func (tu *TeamsUseCase) GetTeams(ctx *gin.Context, collection *mongo.Collection, offset int, limit int) ([]domain.TeamResponse, error) {
	teams, err := tu.teamsRepository.GetTeams(ctx, collection, limit, offset)
	
	if err!=nil && teams != nil {
		return teams, nil
	}
	
	var teamResponses []domain.TeamResponse
	for _, team := range teams {
		teamResponse := domain.TeamResponse{
			Area:           team.Area,
			ID:             team.ID,
			Name:           team.Name,
			ShortName:      team.ShortName,
			Tla:            team.Tla,
			CrestUrl:       team.CrestUrl,
			Address:        team.Address,
			Website:        team.Website,
			Founded:        team.Founded,
			ClubColors:     team.ClubColors,
			Venue:          team.Venue,
			Competition:    team.Competition,
			Squad:          team.Squad,
			LastUpdated:    team.LastUpdated,
		}
		teamResponses = append(teamResponses, teamResponse)
	}

	return teamResponses, nil
}