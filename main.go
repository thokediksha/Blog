package main

import (
	// "fmt"
	"os"
	routers "Blog/routes"
	models "Blog/models"
	"github.com/joho/godotenv"

)

func main(){
	godotenv.Load()          // Load env variables
    models.ConnectDataBase() // load dbe.
	r := routers.InitialzeRoutes()
	
	port := os.Getenv("SERVER_PORT")

    if port == "" {
        port = "8002"
    }

	r.Run(":" + port)
}