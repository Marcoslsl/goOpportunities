package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Router
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)

	// Run the server
	r.Run(":8080")
}
