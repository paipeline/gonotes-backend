// migrate/migrate.go
package main

import (
	"goauth/initializers"
	"goauth/models"
)

func init(){
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.User{}) // automigrate
}
