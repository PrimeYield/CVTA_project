package database

import (
	"context"
	"exercise/pkg/setting"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var databaseName string

func MongodbJoin(databaseSEtting *setting.DatabaseSetting) error {
	clientOptions := options.Client().ApplyURI("mongodb://" + databaseSEtting.MongodbHost + ":" + databaseSEtting.MongodbPort)

	var err error
	Client ,err = mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		log.Printf("package database: connect mongodb fail: %s", err.Error())
		return err
	}
	databaseName = databaseSEtting.Mongodb_db
	log.Println("Connected to MongoDB!")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized. Call MongodbJoin first.")
	}
	return Client.Database(databaseName).Collection(collectionName)
}