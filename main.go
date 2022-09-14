package main

import (
	"book_management/config"
	"book_management/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authcontroller controller.AuthController = controller.NewAuthController()
)

func main() {

	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	//auth routes
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", authcontroller.Login)
		authRoutes.POST("/register", authcontroller.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
