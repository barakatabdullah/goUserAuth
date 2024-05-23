package main

import (
	"example/auth/initializers"
	"example/auth/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	
     initializers.DB.AutoMigrate(&models.User{})
}