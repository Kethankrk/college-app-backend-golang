package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var AttendanceDatabase *mongo.Database
var ClassDatabase *mongo.Database

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	option := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	var err error
	MongoClient, err = mongo.Connect(ctx, option)

	if err != nil {
		log.Fatal("Failed to connect to mongodb")
	} else {
		log.Println("Successfuly connected to mongodb")
	}

	AttendanceDatabase = MongoClient.Database(os.Getenv("ATTENDANCE"))
}
