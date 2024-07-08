package database

import (
    "context"
    "log"
    "sync"
    "time"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "os"
)

var client *mongo.Client
var dbName string
var once sync.Once

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    dbName = os.Getenv("DB_NAME")
}

func Connect() {
    once.Do(func() {
        dbUrl := os.Getenv("DB_URL")
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        var err error
        client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
        if err != nil {
            log.Fatalf("Failed to connect to database: %v", err)
        }

        err = client.Ping(ctx, readpref.Primary())
        if err != nil {
            log.Fatalf("Failed to ping database: %v", err)
        }

        log.Println("Connected to MongoDB!")
    })
}

func GetDatabase() *mongo.Database {
    if client == nil {
        log.Fatal("No client found, please connect to the database first")
    }
    return client.Database(dbName)
}