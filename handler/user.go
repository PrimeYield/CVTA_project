package handler

import (
	"exercise/pkg/logic"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// func SingupHandler(c gin.Context) {

// }

func GetAllRecordHandler(c *gin.Context){
	username := c.Param("username")

	var search bson.M
	_ = c.ShouldBindJSON(&search)

	result,err := logic.GetAllRecord(username)
	if err != nil {
		log.Printf("package handler user GetAllRecord() error:%v",err)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusOK,gin.H{
		"message": "successful",
		"result": result,
	})
}