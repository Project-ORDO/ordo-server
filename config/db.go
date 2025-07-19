package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	uri := os.Getenv("MONGODB_URI") 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connect error:", err)
	}

	// Ping to check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	log.Println("Connected to MongoDB Atlas")
	Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
	dbName := os.Getenv("MONGODB_DB_NAME") 
	return Client.Database(dbName).Collection(collectionName)
}