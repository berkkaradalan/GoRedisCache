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

func NewAreasRouter(env *bootstrap.Env, timeout time.Duration, mongodb *mongo.Collection, redisClient *redis.Client,r *gin.RouterGroup){
	//? areas repository
	ar := repository.NewAreasRepository(mongodb)
	
	//? areas controller
	ac := &controller.AreaController{
		AreaUseCase: usecase.NewAreaUseCase(ar, timeout, env),
		Env: env,
	}

	areasGroup := r.Group("/areas")
	{
		areasGroup.GET("/:id", func(c *gin.Context) {ac.GetAreaById(c, mongodb, redisClient)})
	}
}