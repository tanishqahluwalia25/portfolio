package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDb() *mongo.Client {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error while opening the env file")
	}

	mongoDb := os.Getenv("MONGO_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conncted to mongo")

	return client

}

var Client = connectToDb()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	return client.Database("cluster0").Collection(collectionName)

}
