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
var SquatRecord models.Squat

//when user click "pushup" buttom
//要回傳id給前端
func CreateSquatRecord(username string) (primitive.ObjectID,error) {
	collection := database.GetCollection("squat")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	SquatRecord = models.NewSquat()
	SquatRecord.Username = username
	result, err := collection.InsertOne(ctx,SquatRecord)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("package logic failed: insert squat: %v", err)
	}
	return result.InsertedID.(primitive.ObjectID),nil
}

func GetSquatRecordByID(id primitive.ObjectID) (*models.Squat,error){
	colletction := database.GetCollection("squat")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	var squat models.Squat
	filter := bson.M{"_id": id}
	err := colletction.FindOne(ctx,filter).Decode(&squat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil,nil
		}
		log.Printf("package logic Get Squat Record error: %v",err)
		return nil, fmt.Errorf("package logic Get Squat Record error: %v",err)
	}
	return &squat,nil
}

func UpdateSquatRecord(id primitive.ObjectID,updates bson.M) error {
	collection := database.GetCollection("squat")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	// filter := bson.M{"_id": id}

	record, err := GetSquatRecordByID(id)
	if err != nil {
		log.Printf("package logic Update Squat Record error from Get Squat Record: %v",err)
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
		log.Printf("package logic UpdateByID(squat) err: %v",err)
		return err
	}
	return nil
}