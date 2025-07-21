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

// PushupRecord := models.NewPushup()
var PushupRecord models.Pushup

//when user click "pushup" buttom
//要回傳id給前端
func CreatePushupRecord(username string) (primitive.ObjectID,error) {
	collection := database.GetCollection("pushup")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	PushupRecord = models.NewPushup()
	PushupRecord.Username = username
	result, err := collection.InsertOne(ctx,PushupRecord)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("package logic failed: insert pushup: %v", err)
	}
	return result.InsertedID.(primitive.ObjectID),nil
}

func GetPushupRecordByID(id primitive.ObjectID) (*models.Pushup,error){
	colletction := database.GetCollection("pushup")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	var pushup models.Pushup
	filter := bson.M{"_id": id}
	err := colletction.FindOne(ctx,filter).Decode(&pushup)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil,nil
		}
		log.Printf("package logic Get Pushup Record error: %v",err)
		return nil, fmt.Errorf("package logic Get Pushup Record error: %v",err)
	}
	return &pushup,nil
}

func UpdatePushupRecord(id primitive.ObjectID,updates bson.M) error {
	collection := database.GetCollection("pushup")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	// filter := bson.M{"_id": id}

	record, err := GetPushupRecordByID(id)
	if err != nil {
		log.Printf("package logic Update Pushup Record error from Get Pushup Record: %v",err)
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
		log.Printf("package logic UpdateByID(pushup) err: %v",err)
		return err
	}
	return nil
}