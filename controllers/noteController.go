package controllers

import (
	"goauth/initializers"
	"goauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateNote(c *gin.Context) {
	// 从上下文中获取当前用户
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}
	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户类型断言失败"})
		return
	}

	// 绑定请求数据
	var input struct {
		Title   string `json:"title" binding:"required,max=255"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建笔记
	note := models.Note{
		Title:   input.Title,
		Content: input.Content,
		UserID:  currentUser.ID,
	}

	result := initializers.DB.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建笔记失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "笔记创建成功",
		"note":    note,
	})
}

func ListNotes(c *gin.Context) {
	var notes []models.Note
	initializers.DB.Find(&notes)
	c.JSON(http.StatusOK, gin.H{"data": notes})
 }


func UpdateNote(c *gin.Context) { /* ... */ }
func DeleteNote(c *gin.Context) { /* ... */ }

func GetNote(c *gin.Context) { /* ... */ }
func GetAllNotesByUser(c *gin.Context) { /* ... */ }


