package bootstrap

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(env *Env) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.MONGODB_USERNAME, env.MONGODB_PASSWORD, env.MONGODB_URL, env.MONGODB_PORT)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	  panic(err)
	}
	
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
 		panic(err)
	}

	log.Println("Connected to MongoDB on url: ", env.MONGODB_URL, " port: ", env.MONGODB_PORT)
	return client
}