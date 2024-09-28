package controllers

import (
	"goauth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	GetUserProfile(c *gin.Context)
}

type authController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) AuthController {
	return &authController{
		db: db,
	}
}

// CreateUser godoc
// @Summary 创建新用户
// @Tags 认证
// @Accept json
// @Produce json
// @Param user body models.AuthInput true "用户注册信息"
// @Success 200 {object} models.User
// @Failure 400 {object} string
// @Router /api/v1/auth/signup [post]
func (ac *authController) CreateUser(c *gin.Context) {
	var authInput models.AuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	ac.db.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	ac.db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ac *authController) Login(c *gin.Context) {
	var authInput models.AuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	ac.db.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *authController) GetUserProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}

	userResponse := gin.H{
		"id":        userModel.ID,
		"username":  userModel.Username,
		"createdAt": userModel.CreatedAt,
		"updatedAt": userModel.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}
