package database

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"github.com/joho/godotenv"
)

var dbName string
var dbUrl string

func init(){
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	dbName = os.Getenv("DB_NAME")
	dbUrl = os.Getenv("DB_URL")
}

func Connect() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

