package bootstrap

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(env *Env) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_URL+":"+env.REDIS_PORT,
		Password: env.REDIS_PASSWORD,
		DB:       0,  
	})

	_, err:=client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
		panic(err)
	}
 
	log.Println("Connected to Redis on url  : ", env.REDIS_URL, " port: ", env.REDIS_PORT)
	return client

}