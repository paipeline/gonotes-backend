// migrate/migrate.go
package main

import (
	"goauth/initializers"
	"goauth/models"
	"log"
)

func init(){
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main(){
	
	err := initializers.DB.AutoMigrate(&models.User{}) // automigrate
	if err != nil {
		log.Fatal("Error migrating models:", err)
	}
	log.Println("Migration completed")
	
}
