package logic

import (
	"context"
	"exercise/database"
	"exercise/models"
	jwt "exercise/pkg/JWT"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var User models.User

func CreateNewUser(username,password string,height,weight float64)(primitive.ObjectID,error){
	// if getUsername(username) {
	// 	log.Printf("package logic Username is already exists")
	// 	return primitive.NilObjectID,fmt.Errorf("package logic Username is already exists")
	// }
	
	collection := database.GetCollection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	User = models.NewUser(username,password)
	User.Height = height
	User.Weight = weight
	result, err := collection.InsertOne(ctx,User)
	if err != nil {
		return primitive.NilObjectID,fmt.Errorf("package logic failed: insert user: %v",err)
	}
	return result.InsertedID.(primitive.ObjectID),nil
}

func Login(username,password string) (string,error) {
	user,err := GetUserByUsername(username)
	if err != nil {
		return "",err
	}
	if user == nil {
		return "",fmt.Errorf("%s is not exists" ,username)
	}
	if user.Password != password {
		log.Print("wrong password")
		return "", fmt.Errorf("wrong password:%s",password)
	}
	token,err := jwt.GenerateToken(username)
	if err != nil {
		return "",fmt.Errorf("package logic Login user fail:%v",err)
	}
	return token,nil
}

func getUsername(username string) bool {
	collection := database.GetCollection("user")
	ctx, cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	_, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func GetUserByUsername(username string) (*models.User,error){
	collection := database.GetCollection("user")
	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()

	filter := bson.M{"username":username}
	cursor := collection.FindOne(ctx,filter)

	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil,fmt.Errorf("username : %s is not exists",username)
		}
		return nil,fmt.Errorf("package logic GetUserByUsername bad request for %s",username)
	}
	return &user,nil
}

//功能：
//login後要獲得所有user的運動記錄

func GetAllRecord(username string)(*models.UserRecord,error){
	if !getUsername(username) {
		log.Printf("package logic user: user is not exist")
		return nil,fmt.Errorf("package logic user: user is not exist")
	}
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