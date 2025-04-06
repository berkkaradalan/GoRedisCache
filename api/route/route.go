package route

import (
	"time"
	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, mongoClient *mongo.Client, redisclient *redis.Client,r *gin.Engine){
	public := r.Group("/api/")
	collections := bootstrap.GetCollections(mongoClient)
	NewAreasRouter(env, timeout, collections.Areas, redisclient, public)
	NewTeamsRouter(env, timeout, collections.Teams, redisclient, public)
	NewPlayerRouter(env, timeout, collections.Players, redisclient, public)
	NewMatchesRouter(env, &timeout, collections.Matches, public)
}