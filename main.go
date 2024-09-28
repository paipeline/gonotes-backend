// @title GoAuth API
// @version 1.0
// @description 用户认证和笔记管理 API
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"goauth/api"
	"goauth/controllers"
	"goauth/initializers"
	"goauth/repositories"
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

	// 初始化数据库连接
	db := initializers.DB

	// 初始化存储库
	documentRepo := repositories.NewDocumentRepository(db)
	authController := controllers.NewAuthController(db)

	// 初始化 API
	userAPI := api.NewUserAPI(authController, documentRepo)

	// 设置路由
	userAPI.SetupRoutes(router)

	router.Run(":8080")
}
