package main

import (
	"goauth/controllers"
	"goauth/initializers"

	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main(){
	router := gin.Default()
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile",controllers.GetUserProfile)
	router.Run(":8080")
}
