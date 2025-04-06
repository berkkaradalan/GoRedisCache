package route

import (
	"github.com/berkkaradalan/GoRedisCache/api/controller"
	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/repository"
	"github.com/berkkaradalan/GoRedisCache/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTeamsRouter(env *bootstrap.Env, timeout time.Duration, mongodb *mongo.Collection, redisClient *redis.Client,r *gin.RouterGroup){
	tr := repository.NewTeamsRepository(mongodb)
	
	tc := &controller.TeamsController{
		TeamsUseCase: usecase.NewTeamsUseCase(tr, timeout, env),
		Env: env,
	}

	teamsGroup := r.Group("/teams")
	{
		teamsGroup.GET("/:id", func(c *gin.Context) {tc.GetTeamById(c, mongodb, redisClient)})
		teamsGroup.GET("", func(c *gin.Context) {tc.GetTeams(c, mongodb)})
	}
}