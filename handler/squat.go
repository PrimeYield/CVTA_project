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

func CreateSquatRecordHandler(c *gin.Context) {
	squatRecord := models.Squat{}
	if err := c.ShouldBindJSON(&squatRecord); err != nil {
		log.Printf("handle BindJson(squat) error: %v", err)
		return
	}

	recordID, err := logic.CreateSquatRecord(squatRecord.Username)
	if err != nil {
		log.Printf("handle CreateSquatRecord error: %v", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "New squat record is created",
		"id":       recordID.Hex(),
		"username": squatRecord.Username,
		"count":    squatRecord.Count,
		"calorie":  squatRecord.Calorie,
		"date":     squatRecord.Date,
		"start":    squatRecord.Start,
	})
}

func UpdateSquatRecordHandler(c *gin.Context) {
	idStr := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("package main ObjectIDFromHex error: %v", err)
		return
	}

	var updates bson.M
	_ = c.ShouldBindJSON(&updates)

	err = logic.UpdateSquatRecord(objID, updates)
	if err != nil {
		log.Printf("package main UpdateRecord(squat) error: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "good action",
		"id":      objID.Hex(),
	})
}