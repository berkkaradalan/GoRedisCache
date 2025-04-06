package bootstrap

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/redis/go-redis/v9"
)
type Application struct {
	Env *Env
	MongoDB *mongo.Client
	Redis   *redis.Client
}

func App() Application{
	app := Application{}
	app.Env = NewEnv()
	app.MongoDB = ConnectToMongoDB(app.Env)
	app.Redis = ConnectToRedis(app.Env)
	return app
}