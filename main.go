package main

import (
	"github.com/MohamedOuhami/TeamGamifierWithGo/controllers"
	"github.com/MohamedOuhami/TeamGamifierWithGo/initializers"
	"github.com/MohamedOuhami/TeamGamifierWithGo/middleware"
	"github.com/gin-gonic/gin"
)

// In the Init function, we're going to initialize the env variables
func init() {
	// Initialize the env vars
	initializers.LoadEnv()

	// Connect to the database
	initializers.ConnectToDb()

	// Sync with the database
	initializers.SyncDatabase()
}

func main() {

	// Started the GIN server
	r := gin.Default()

	// User Endpoints
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.IsAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:3000

}
