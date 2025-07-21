package handler

import (
	"exercise/models"
	"exercise/pkg/logic"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//07.21計畫新增Token攜帶login後的username
func CreatePushupRecordHandler(c *gin.Context) {
	pushupRecord := models.Pushup{}
	if err := c.ShouldBindJSON(&pushupRecord); err != nil {
		log.Printf("handle BindJson error: %v", err)
		return
	}

	recordID, err := logic.CreatePushupRecord(pushupRecord.Username)
	if err != nil {
		log.Printf("handle CreateRecord error: %v", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "New pushup record is created",
		"id":       recordID.Hex(),
		"username": pushupRecord.Username,
		"count":    pushupRecord.Count,
		"calorie":  pushupRecord.Calorie,
		"date":     pushupRecord.Date,
		"start":    pushupRecord.Start,
	})
}

func UpdatePushupRecordHandler (c *gin.Context){
	idStr := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("package main ObjectIDFromHex error: %v",err)
		return
	}

	var updates bson.M
	_ = c.ShouldBindJSON(&updates)
	

	err = logic.UpdatePushupRecord(objID,updates)
	if err != nil {
		log.Printf("package main UpdateRecord error: %v",err)
		return
	}
	newRecord,err := logic.GetPushupRecordByID(objID)
	if err != nil {
		log.Printf("package handler Update Get error:%v",err)
	}
	c.JSON(http.StatusOK,gin.H{
		// "message": "good action",
		// "id": objID.Hex(),
		"count": newRecord.Count,
	})
}