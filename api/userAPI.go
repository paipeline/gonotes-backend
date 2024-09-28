package api

import (
	"goauth/controllers"
	"goauth/middlewares"
	"goauth/repositories"

	"github.com/gin-gonic/gin"
)

// API 接口定义了所有 API 的行为
type API interface {
	SetupRoutes(router *gin.Engine)
}

// UserAPI 实现了 API 接口
type UserAPI struct {
	authController     controllers.AuthController
	documentRepository repositories.DocumentRepository
}

// NewUserAPI 创建一个新的 UserAPI 实例
func NewUserAPI(
	authController controllers.AuthController,
	documentRepository repositories.DocumentRepository,
) *UserAPI {
	return &UserAPI{
		authController:     authController,
		documentRepository: documentRepository,
	}
}

// SetupRoutes 设置所有的路由
func (api *UserAPI) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	api.setupAuthRoutes(v1)
	api.setupUserRoutes(v1)
	// api.setupDocumentRoutes(router)
}

// setupAuthRoutes 设置认证相关的路由
func (api *UserAPI) setupAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/signup", api.authController.CreateUser)
		auth.POST("/login", api.authController.Login)
	}
}

// setupUserRoutes 设置用户相关的路由
func (api *UserAPI) setupUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	user.Use(middlewares.CheckAuth)
	{
		user.GET("/profile", api.authController.GetUserProfile)
		// api.setupDocumentRoutes(user)
	}
}

// 如果需要 setupDocumentRoutes 方法，请添加如下定义
// func (api *UserAPI) setupDocumentRoutes(router *gin.Engine) {
//     // 设置文档相关路由
// }
