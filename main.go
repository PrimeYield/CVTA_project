package main

import (
	"context"
	"exercise/database"
	"exercise/global"
	"exercise/handler"
	"exercise/pkg/setting"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwa"
)

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		panic(err)
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		log.Printf("package main setupSetting \"Server\" error: %v",err)
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		log.Printf("package main setupSetting \"Database\" error: %v",err)
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		log.Printf("package main setupSetting \"JWT\" error: %v",err)
		return err
	}
	global.JWTSetting.Algorithm = jwa.HS256
	return nil
}

func main(){
	setupSetting()
	err := database.MongodbJoin(&global.DatabaseSetting)
	if err != nil {
		panic("Connect to db failed:" + err.Error())
	}
	defer func() {
		err := database.Client.Disconnect(context.Background())
		if err != nil {
			log.Printf("package main Disconnect DB error: %v",err)
			// return
		}
	}()

	r := gin.Default()
	port := global.ServerSetting.Port
	
	r.Static("public", "./public")
	r.GET("/", func(c *gin.Context) {
	c.File("./public/index.html")
	})

	user := r.Group("/user")
	{
		user.POST("/singup",handler.SingupHandler)
		user.POST("/login",handler.LoginHandler)
		user.GET("/records/:username", handler.GetAllRecordHandler)
	}

	pushup := r.Group("/pushup")
	{
		pushup.POST("/start", handler.CreatePushupRecordHandler)
		pushup.PATCH("/update/:id",handler.UpdatePushupRecordHandler)
	}
	situp := r.Group("/situp")
	{
		situp.POST("/start", handler.CreateSitupRecordHandler)
		situp.PATCH("/update/:id",handler.UpdateSitupRecordHandler)

	}
	squat := r.Group("/squat")
	{
		squat.POST("/start", handler.CreateSquatRecordHandler)
		squat.PATCH("/update/:id",handler.UpdateSquatRecordHandler)

	}
	r.Run(":" + port)
}