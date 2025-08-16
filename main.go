package main

import (
	"github.com/gin-gonic/gin"
	"go-monitoring-service/controllers/dashboardController"
	"go-monitoring-service/controllers/monitorController"
	"go-monitoring-service/controllers/userController"
	"go-monitoring-service/services"
	"go-monitoring-service/initializers"
	middleware "go-monitoring-service/middleware/auth"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	go services.MonitorScheduler()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.GET("/add-monitor", middleware.RequireAuth, func(c *gin.Context) {
		c.HTML(200, "add-monitor.html", nil)
	})
	r.POST("/signup", userController.Signup)
	r.POST("/login", userController.Login)
	r.GET("/logout", userController.Logout)

	r.POST("/add-monitor", middleware.RequireAuth, monitorController.AddMonitor)
	r.GET("/dashboard", middleware.RequireAuth, dashboardController.Dashboard)

	err := r.Run()
	if err != nil {
		return
	}
}
