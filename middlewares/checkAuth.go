package middlewares

import (
	"fmt"
	"goauth/initializers"
	"goauth/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)




func CheckAuth(c *gin.Context){
	authHeader:=c.GetHeader("authorization")
	if authHeader==""{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Authorzation header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken != 2) || authToken[0] != "Bearer"{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1];
	token, err := jwt.Parse(tokenSring, func(token *jwt.Token) (interface{} error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("Unexpected singing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.JSON(http.StatusUnauthorized, gin.H{"error":"token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	initializers.DB.Where("ID=?", claims["ID"]).Find(&user)

	if user.ID == 0{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)

	c.Next()
}