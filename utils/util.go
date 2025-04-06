package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetInt(m bson.M, key string) int {
	if val, ok := m[key].(float64); ok {
		return int(val)
	}
	return 0
}

func GetString(m bson.M, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}

func GetStringPointer(m bson.M, key string) *string {
	if val, ok := m[key].(string); ok {
		return &val
	}
	return nil
}

func JSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.Encode(obj)
}

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func MigrateDB(db *mongo.Client, db_name string) {
	collections, err := db.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to list collections: %v", err)
	}
	database := db.Database(db_name)

	if !Contains(collections, "areas") {
		err = database.CreateCollection(context.TODO(), "areas")
		if err != nil {
			log.Fatalf("Error while creating collection: %v", err)
		}

		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		}

		_, err = database.Collection("areas").Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Fatalf("Error creating index on areas collection: %v", err)
		}
	}

	if !Contains(collections, "teams") {
		err = database.CreateCollection(context.TODO(), "teams")
		if err != nil {
			log.Fatalf("Error while creating collection: %v", err)
		}

		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		}

		_, err = database.Collection("teams").Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Fatalf("Error creating index on teams collection: %v", err)
		}
	}

	if !Contains(collections, "players") {
		err = database.CreateCollection(context.TODO(), "players")
		if err != nil {
			log.Fatalf("Error while creating collection: %v", err)
		}

		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: "id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		}

		_, err = database.Collection("players").Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Fatalf("Error creating index on players collection: %v", err)
		}
	}
}

func RequestToExternalApi(api_key string, url string) (bson.M, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error while creating request: %v", err)
		return nil, err
	}

	req.Header.Set("X-Auth-Token", api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error while sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error while reading response body: %v", err)
		return nil, err
	}

	var result bson.M
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Error while unmarshalling JSON: %v", err)
		return nil, err
	}

	return result, nil
}