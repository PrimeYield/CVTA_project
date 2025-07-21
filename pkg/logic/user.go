package logic

import (
	"context"
	"exercise/database"
	"exercise/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var User models.User

func CreateNewUser(username,password string)(primitive.ObjectID,error){
	collection := database.GetCollection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	User = models.NewUser(username,password)
	result, err := collection.InsertOne(ctx,User)
	if err != nil {
		return primitive.NilObjectID,fmt.Errorf("package logic failed: insert user: %v",err)
	}
	return result.InsertedID.(primitive.ObjectID),nil
}

//功能：
//login後要獲得所有user的運動記錄

func GetAllRecord(username string)(*models.UserRecord,error){
	records := models.UserRecord{}
	pushupCollection := database.GetCollection("pushup")
	pushupCursor,err := pushupCollection.Find(context.TODO(),bson.M{"username":username})
	if err != nil {
		return nil,fmt.Errorf("package logic failed: GetAllRecord(pushup) error:%v",err)
	}
	defer pushupCursor.Close(context.TODO())

	if err = pushupCursor.All(context.TODO(),&records.Pushups); err != nil {
		return nil,fmt.Errorf("package logic failed: pushupCursor.All error: %v",err)
	}

	situpCollection := database.GetCollection("situp")
	situpCursor,err := situpCollection.Find(context.TODO(),bson.M{"username":username})
	if err != nil {
		return nil,fmt.Errorf("package logic failed GetAllRecord(situp) error:%v", err)
	}
	defer situpCursor.Close(context.TODO())
	
	if err = situpCursor.All(context.TODO(),&records.Situps); err != nil {
		return nil,fmt.Errorf("package logic failed: situpCursor.All error:%v",err)
	}

	squatCollection := database.GetCollection("squat")
	squatCursor,err := squatCollection.Find(context.TODO(),bson.M{"username":username})
	if err != nil {
		return nil,fmt.Errorf("package logic failed GetAllRecord(squat) error:%v",err)
	}
	defer squatCursor.Close(context.TODO())

	if err = squatCursor.All(context.TODO(),&records.Squats);err != nil {
		return nil, fmt.Errorf("package logic failed: squatCursor.All error:%v",err)
	}
	return &records,nil
}