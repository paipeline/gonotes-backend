// controllers/authController.go
package controllers

import (
	"goauth/initializers"
	"goauth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(c *gin.Context){
	var authInput models.AuthInput 
	if err:= c.ShouldBindJSON(&authInput); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return;
	}


	// check if found
	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != 0{
		c.JSON(http.StatusBadRequest, gin.H{"error":"username already exists"})

	}
	// hash password
	passwordHash, err:= bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}


	// store user
	user := models.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	// main logic of storage 
	initializers.DB.Create(&user)
	c.JSON(http.StatusOK,gin.H{"data":user})


}



func Login(c *gin.Context){
	var authInput models.AuthInput // here the user input
	if err := c.ShouldBindJSON(&authInput); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	var userFound models.User; // compare the input to the database username
	initializers.DB.Where("username=?",authInput.Username).Find(&userFound)

	if userFound.ID == 0 { // not found
		c.JSON(http.StatusBadRequest,gin.H{"error":"User not found"})
	}
	// user ID found but not correct password
	if err:= bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{	
		"id": userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token":token,
	})

}



func GetUserProfile(c* gin.Context){
	user,_ := c.Get("currentUser");
	c.JSON(200,gin.H{
		"user":user,
	})

}

