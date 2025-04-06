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
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPlayerUseCase(pr repository.PlayersRepository, timeout time.Duration, env *bootstrap.Env) *PlayerUseCase {
	return &PlayerUseCase{
		playerRepository: 			pr,
		contextTimeout: 			timeout,
		Env: 						env,
	}
}

type PlayerUseCase struct{
	playerRepository 				repository.PlayersRepository
	contextTimeout 					time.Duration
	Env 							*bootstrap.Env
}

func (pu *PlayerUseCase) GetPlayerById(ctx *gin.Context, collection *mongo.Collection, redisClient *redis.Client,player_id string) (domain.PlayerResponse, error) {
	cacheKey := "player:" + player_id
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedPlayer domain.PlayerResponse
		if jsonErr := json.Unmarshal([]byte(cachedData), &cachedPlayer); jsonErr == nil {
			return cachedPlayer, nil
		}
	}

	player, err := pu.playerRepository.GetPlayerById(ctx, collection, player_id)
	if err == nil && player != nil {
		return *player, nil
	}

	resp, err := utils.RequestToExternalApi(pu.Env.FOOTBALL_DATA_API_KEY, pu.Env.FOOTBALL_DATA_PLAYER_API_URL+player_id)
	if err != nil {
		return domain.PlayerResponse{}, fmt.Errorf("no player found with id: %s", player_id)
	}

	if errMsg, exists := resp["error"]; exists {
		if errCode, ok := errMsg.(float64); ok && int(errCode) == 404 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Player not found in external API"})
			return domain.PlayerResponse{}, fmt.Errorf("player with id %s not found", player_id)
		}
	}

	getInt := utils.GetInt
	getString := utils.GetString
	getStringPointer := utils.GetStringPointer
	
	var team *domain.PlayerTeam
	if teamData, ok := resp["currentTeam"].(map[string]interface{}); ok {

		var contract *domain.PlayerContract
		if contractData, exists := teamData["contract"].(map[string]interface{}); exists {
			contract = &domain.PlayerContract{
				Start: getString(contractData, "start"),
				Until: getString(contractData, "until"),
			}
		}

		var competitions []*domain.PlayerTeamRunninCompetitions
		if competitionsData, exists := teamData["runningCompetitions"].([]interface{}); exists {
			for _, comp := range competitionsData {
				if compMap, ok := comp.(map[string]interface{}); ok {
					competitions = append(competitions, &domain.PlayerTeamRunninCompetitions{
						ID:        getInt(compMap, "id"),
						Name:      getString(compMap, "name"),
						Code:      getString(compMap, "code"),
						Type:      getString(compMap, "type"),
						EmblemUrl: getStringPointer(compMap, "emblem"),
					})
				}
			}
		}

		areaData := teamData["area"].(map[string]interface{})
		team = &domain.PlayerTeam{
			Area: &domain.PlayerTeamArea{
				ID:          getInt(areaData, "id"),
				Name:        getString(areaData, "name"),
				CountryCode: getString(areaData, "code"),
				FlagUrl:     getStringPointer(areaData, "flag"),
			},
			ID:           	 getInt(teamData, "id"),
			Name:         	 getString(teamData, "name"),
			ShortName:    	 getString(teamData, "shortName"),
			Tla:          	 getString(teamData, "tla"),
			CrestUrl:     	 getStringPointer(teamData, "crest"),
			Address:      	 getStringPointer(teamData, "address"),
			Website:      	 getStringPointer(teamData, "website"),
			Founded:      	 getInt(teamData, "founded"),
			ClubColors:   	 getString(teamData, "clubColors"),
			Venue:        	 getString(teamData, "venue"),
			Contract:     	 contract,
			Competitions: 	 competitions,
		}
	}

	var lastUpdated *string
	if val, ok := resp["lastUpdated"].(string); ok {
		lastUpdated = &val
	}

	convertedPlayer := domain.PlayerResponse{
		ID:          		getInt(resp, "id"),
		Name:        		getString(resp, "name"),
		FirstName:   		getString(resp, "firstName"),
		LastName:    		getString(resp, "lastName"),
		DateOfBirth: 		getString(resp, "dateOfBirth"),
		Nationality: 		getString(resp, "nationality"),
		Section:     		getString(resp, "section"),
		Position:    		getString(resp, "position"),
		ShirtNumber: 		getInt(resp, "shirtNumber"),
		LastUpdated: 		lastUpdated,
		CurrentTeam:        team,
	}

	_, err = collection.InsertOne(ctx, convertedPlayer)
	if err != nil {
		log.Printf("MongoDB insert error: %v", err)
	}

	cachePlayer(ctx, redisClient, cacheKey, &convertedPlayer)
	return convertedPlayer, nil
}


func cachePlayer(ctx context.Context, redisClient *redis.Client, key string, player *domain.PlayerResponse) {
	playerJSON, err := json.Marshal(player)
	if err != nil {
		log.Printf("Failed to marshal player data: %v", err)
		return
	}
	status := redisClient.Set(ctx, key, playerJSON, time.Hour)
	if status.Err() != nil {
        log.Printf("Redis set error: %v", status.Err())
    } else {
        log.Printf("Successfully cached player: %s", key)
    }
}