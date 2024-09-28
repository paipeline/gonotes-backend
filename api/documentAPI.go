package api

import (
	"goauth/controllers"
	"goauth/models"
	"goauth/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DocumentAPI 结构体用于处理文档相关的 API 请求
type DocumentAPI struct {
	authController     controllers.AuthController
	documentRepository repositories.DocumentRepository
}

// NewDocumentAPI 创建一个新的 DocumentAPI 实例
func NewDocumentAPI(
	authController controllers.AuthController,
	documentRepository repositories.DocumentRepository) *DocumentAPI {

	return &DocumentAPI{
		authController:     authController,
		documentRepository: documentRepository,
	}
}

// SetupRoutes 设置文档相关的路由
func (api *DocumentAPI) SetupRoutes(router *gin.Engine) {
	documents := router.Group("/documents")
	{
		documents.POST("", api.CreateDocument)
		documents.GET("/:id", api.GetDocument)
		documents.PUT("/:id", api.UpdateDocument)
		documents.DELETE("/:id", api.DeleteDocument)
	}
}

// CreateDocument 处理创建新文档的请求
func (api *DocumentAPI) CreateDocument(c *gin.Context) {
	var document models.Document
	// 从请求体中解析 JSON 数据
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用仓库方法创建文档
	if err := api.documentRepository.CreateDocument(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文档失败"})
		return
	}

	c.JSON(http.StatusOK, document)
}

// GetDocument 处理获取特定文档的请求
func (api *DocumentAPI) GetDocument(c *gin.Context) {
	id := c.Param("id")

	// 调用仓库方法获取文档
	document, err := api.documentRepository.GetDocument(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文档未找到"})
		return
	}

	c.JSON(http.StatusOK, document)
}

// UpdateDocument 处理更新文档的请求
func (api *DocumentAPI) UpdateDocument(c *gin.Context) {
	idStr := c.Param("id")
	// 将字符串 ID 转换为 uint 类型
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var document models.Document
	// 从请求体中解析 JSON 数据
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	document.ID = uint(id) // 假设 Document.ID 是 uint 类型

	// 调用仓库方法更新文档
	if err := api.documentRepository.UpdateDocument(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文档失败"})
		return
	}

	c.JSON(http.StatusOK, document)
}

// DeleteDocument 处理删除文档的请求
func (api *DocumentAPI) DeleteDocument(c *gin.Context) {
	id := c.Param("id")

	// 调用仓库方法删除文档
	if err := api.documentRepository.DeleteDocument(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文档失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文档已成功删除"})
}
