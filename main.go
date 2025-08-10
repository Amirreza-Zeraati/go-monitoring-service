package main

import (
	"github.com/gin-gonic/gin"
	"go-monitoring-service/initializers"
	middleware "go-monitoring-service/middleware/auth"
	userServices "go-monitoring-service/services/userService"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/signup", userServices.Signup)
	r.POST("/login", userServices.Login)

	r.GET("/dashboard", middleware.RequireAuth, userServices.Dashboard)

	err := r.Run()
	if err != nil {
		return
	}
}
