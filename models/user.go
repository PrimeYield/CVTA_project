package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	CreateDate string `bson:"create_date" json:"create_date"`
	LoginTimes int `bson:"login_times" json:"login_times"`
}

func NewUser(username, password string) User{
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")
	return User{Username: username,Password: password,CreateDate: date,LoginTimes: 1}
}

func (u *User) UpdateLoginTimes (){
	u.LoginTimes++
}

type UserRecord struct {
	Pushups []Pushup `json:"pushups"`
	Situps  []Situp  `json:"situps"`
	Squats  []Squat  `json:"squats"`
}