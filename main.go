package main

import (
	routers "music/web-music-microservice/app"
	"os"
)

func main() {
	router := routers.InitializeRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
