package main

import (
	"goauth/controllers"
	"goauth/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "healthy service")
	})

	// api control
	v1 := router.Group("/api/v1")

	// auth related operations
	auth := v1.Group("/auth")
	{
		auth.POST("/signup", controllers.CreateUser)
		auth.POST("/login", controllers.Login)
	}
	// user related operations
	user := v1.Group("/user")
	{
		user.GET("/profile", controllers.GetUserProfile)
		// notes related operations
		notes := v1.Group("/note")
		{
			notes.POST("", controllers.CreateNote)       // example localhost:8080/api/v1/user/note
			notes.GET("/:id", controllers.GetNote)       // example localhost:8080/api/v1/user/note/1
			notes.PUT("/:id", controllers.UpdateNote)    // example localhost:8080/api/v1/user/note/1
			notes.DELETE("/:id", controllers.DeleteNote) // example localhost:8080/api/v1/user/note/1
			notes.GET("/all", controllers.ListNotes)     // example localhost:8080/api/v1/note/all
		}

	}

	router.Run(":8080")
}
