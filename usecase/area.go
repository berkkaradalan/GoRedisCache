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

func NewAreaUseCase(ar repository.AreasRepository, timeout time.Duration, env *bootstrap.Env) *AreaUseCase {
	return &AreaUseCase{
		areaRepository: 			ar,
		contextTimeout: 			timeout,
		Env: 						env,
	}
}

type AreaUseCase struct{
	areaRepository 			repository.AreasRepository
	contextTimeout 			time.Duration
	Env 					*bootstrap.Env
}

func (au *AreaUseCase) GetAreaById(ctx *gin.Context, collection *mongo.Collection, redisClient *redis.Client, area_id string) (domain.AreaResponse, error) {

	cacheKey := "area:" + area_id
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedArea domain.AreaResponse
		if jsonErr := json.Unmarshal([]byte(cachedData), &cachedArea); jsonErr == nil {
			return cachedArea, nil
		}
	}

	area, err := au.areaRepository.GetAreaById(ctx, collection, area_id)
	if err == nil && area != nil {
		return *area, nil
	}

	resp, err := utils.RequestToExternalApi(au.Env.FOOTBALL_DATA_API_KEY, au.Env.FOOTBALL_DATA_AREA_API_URL+area_id)
	if err != nil {
		return domain.AreaResponse{}, fmt.Errorf("no area found with id: %s", area_id)
	}

	if errMsg, exists := resp["error"]; exists {
		if errCode, ok := errMsg.(float64); ok && int(errCode) == 404 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Area not found in external API"})
			return domain.AreaResponse{}, fmt.Errorf("area with id %s not found", area_id)
		}
	}

	if errMsg, exists := resp["errorCode"]; exists {
		if errCode, ok := errMsg.(float64); ok && int(errCode) == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Area Id is not valid"})
			return domain.AreaResponse{}, fmt.Errorf("area with id %s not valid", area_id)
		}
	}

	getInt := func(m bson.M, key string) int {
		if val, ok := m[key].(float64); ok {
			return int(val)
		}
		return 0
	}

	getString := func(m bson.M, key string) string {
		if val, ok := m[key].(string); ok {
			return val
		}
		return ""
	}

	getStringPointer := func(m bson.M, key string) *string {
		if val, ok := m[key].(string); ok {
			return &val
		}
		return nil
	}

	convertedArea := domain.AreaResponse{
		ID:           getInt(resp, "id"),
		Name:         getString(resp, "name"),
		CountryCode:  getString(resp, "code"),
		Flag:         getStringPointer(resp, "flag"),
		ParentAreaID: getInt(resp, "parentAreaId"),
		ParentArea:   getString(resp, "parentArea"),
		ChildAreas:   []*domain.Area{},
	}

	_, err = collection.InsertOne(ctx, convertedArea)
	if err != nil {
		log.Printf("MongoDB insert error: %v", err)
	}

	cacheArea(ctx, redisClient, cacheKey, &convertedArea)
	return convertedArea, nil
}

func cacheArea(ctx context.Context, redisClient *redis.Client, key string, area *domain.AreaResponse) {
	areaJSON, err := json.Marshal(area)
	if err != nil {
		log.Printf("Failed to marshal player data: %v", err)
		return
	}
	status := redisClient.Set(ctx, key, areaJSON, time.Hour)
	if status.Err() != nil {
        log.Printf("Redis set error: %v", status.Err())
    } else {
        log.Printf("Successfully cached player: %s", key)
    }
}