package configs

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"projects/ProjectOneAPI_3/models"
	"time"
)

func ConnectionMDB(dbName, cName string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		fmt.Println("error creating client: ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("error pinging connection: ", err)
	}
	return client.Database(dbName).Collection(cName)
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error .env")
	}

	mongoURI := os.Getenv("MONGOURI")
	return mongoURI
}

func CConnectionMDB(d models.Movie, dbName, cName string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		fmt.Println("error creating client: ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("error pinging connection: ", err)
	}

	if _, err := client.Database(dbName).Collection(cName).InsertOne(context.Background(), d); err != nil {
		fmt.Println("error creating db: ", err)
	}

}
