package route

import (
	"time"
	"github.com/berkkaradalan/GoRedisCache/api/controller"
	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/repository"
	"github.com/berkkaradalan/GoRedisCache/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPlayerRouter(env *bootstrap.Env, timeout time.Duration, mongodb *mongo.Collection, redisClient *redis.Client,r *gin.RouterGroup){
	pr := repository.NewPlayerRepository(mongodb)
	pc := &controller.PlayerController{
		PlayerUseCase: usecase.NewPlayerUseCase(pr, timeout, env),
		Env: env,
	}

	playerGroup := r.Group("/players")
	{
		playerGroup.GET("/:id", func(c *gin.Context) {pc.GetPlayerById(c, mongodb, redisClient)})
	}
}