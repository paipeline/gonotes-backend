package initializers

import (
	"log"

	"github.com/joho/godotenv"
) 
func LoadEnvs(){
	err:=godotenv.Load() // load the env variables
	if err!=nil{
		log.Fatal("Error loading .Env file");
	}
}

