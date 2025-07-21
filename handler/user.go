package handler

import (
	"exercise/global"
	"exercise/models"
	"exercise/pkg/logic"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//post
//signup
func SingupHandler(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("package handler failed : BindJSON(singup) error:%v",err)
		return
	}
	insertedID,err := logic.CreateNewUser(user.Username,user.Password,user.Height,user.Weight)
	if err != nil {
		log.Printf("package handler failed: CreatedNewUser() error:%v",err)
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message": "successful",
		"id": insertedID.Hex(),
		"username": user.Username,
	})
}
//login

type UserInfo struct {
	Username,Password string
}

func LoginHandler(c *gin.Context){
	userInfo := new(UserInfo)
	err := c.ShouldBindJSON(userInfo)
	if err != nil {
		log.Printf("package handler LoginHandler BindJSON error: %v",err)
		return
	}
	token,err := logic.Login(userInfo.Username,userInfo.Password)
	if err != nil {
		log.Printf("package handler LoginHandler Login error: %v",err)
		return
	}
	tokenExpiration := time.Now().Add(global.JWTSetting.Expire)
	maxAge := int(time.Until(tokenExpiration).Seconds())
	c.SetCookie(
		"jwt_token",
		token,
		maxAge,
		"/",
		"localhost",
		true,
		true,
	)
}


//get
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