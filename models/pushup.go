package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pushup struct {
	ID		primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	Username string				`bson:"username" json:"username"`
	Count	float64 			`bson:"count" json:"count"`
	Calorie	float64				`bson:"calorie" json:"calorie"`
	Date	string				`bson:"date" json:"date"`
	Start	string				`bson:"start" json:"start"`
	End		string				`bson:"end" json:"end"`
}

func NewPushup() Pushup {
	now := time.Now()
	date := now.Format("06.01.02")
	start := now.Format("15.04.05")
	return Pushup{Count:0,Calorie:30.7,Date: date,Start: start}
}

func (p *Pushup) TotalCalorie () float64 {
	return p.Calorie * p.Count
}