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

func CreateSitupRecordHandler(c *gin.Context) {
	situpRecord := models.Situp{}
	if err := c.ShouldBindJSON(&situpRecord); err != nil {
		log.Printf("handle BindJson(situp) error: %v", err)
		return
	}

	recordID, err := logic.CreateSitupRecord(situpRecord.Username)
	if err != nil {
		log.Printf("handle CreateSitupRecord error: %v", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "New situp record is created",
		"id":       recordID.Hex(),
		"username": situpRecord.Username,
		"count":    situpRecord.Count,
		"calorie":  situpRecord.Calorie,
		"date":     situpRecord.Date,
		"start":    situpRecord.Start,
	})
}

func UpdateSitupRecordHandler(c *gin.Context) {
	idStr := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("package main Situp ObjectIDFromHex error: %v", err)
		return
	}

	var updates bson.M
	_ = c.ShouldBindJSON(&updates)

	err = logic.UpdateSitupRecord(objID, updates)
	if err != nil {
		log.Printf("package main UpdateRecord(situp) error: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "good action",
		"id":      objID.Hex(),
	})
}