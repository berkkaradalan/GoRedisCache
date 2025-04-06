package route

import (
	"time"

	"github.com/berkkaradalan/GoRedisCache/api/controller"
	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/repository"
	"github.com/berkkaradalan/GoRedisCache/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMatchesRouter(env *bootstrap.Env, timeout *time.Duration, mongodb *mongo.Collection, r *gin.RouterGroup){
	mr := repository.NewMatchesRepository(mongodb)

	mc := &controller.MatchesController{
		MatchesUseCase: usecase.NewMatchesUseCase(mr, *timeout, env),
		Env: env,
	}

	matchesGroup := r.Group("/matches")
	{
		matchesGroup.GET("/team/:id", func(c *gin.Context) {mc.GetTeamMatches(c, mongodb)})
	}
}