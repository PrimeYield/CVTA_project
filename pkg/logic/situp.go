package logic

import (
	"context"
	"exercise/database"
	"exercise/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var SitupRecord models.Situp

//when user click "situp buttom
func CreateSitupRecord(username string) (primitive.ObjectID,error) {
	collection := database.GetCollection("situp")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	SitupRecord = models.NewSitup()
	SitupRecord.Username = username
	result, err := collection.InsertOne(ctx,SitupRecord)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("package logic failed: insert situp: %v", err)
	}
	return result.InsertedID.(primitive.ObjectID),nil
}

func GetSitupRecordByID(id primitive.ObjectID) (*models.Situp,error){
	colletction := database.GetCollection("situp")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	var situp models.Situp
	filter := bson.M{"_id": id}
	err := colletction.FindOne(ctx,filter).Decode(&situp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil,nil
		}
		log.Printf("package logic Get situp Record error: %v",err)
		return nil, fmt.Errorf("package logic Get situp Record error: %v",err)
	}
	return &situp,nil
}

func UpdateSitupRecord(id primitive.ObjectID,updates bson.M) error {
	collection := database.GetCollection("situp")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	// filter := bson.M{"_id": id}

	record, err := GetSitupRecordByID(id)
	if err != nil {
		log.Printf("package logic Update Situp Record error from Get Situp Record: %v",err)
	}
	now := time.Now()
	end := now.Format("15.04.05")
	record.End = end
	// fmt.Println(record.Count)
	record.Count++
	// fmt.Println(record.Count)

	setUpdate := make(bson.M)
	for k,v := range updates {
		setUpdate[k] = v
	}
	setUpdate["end"] = record.End
	setUpdate["count"] = record.Count

	update := bson.M{"$set": setUpdate}

	_ ,err = collection.UpdateByID(ctx,id,update)
	if err != nil {
		log.Printf("package logic UpdateByID(situp) err: %v",err)
		return err
	}
	return nil
}