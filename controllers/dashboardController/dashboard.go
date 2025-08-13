package dashboardController

import (
	"github.com/gin-gonic/gin"
	"go-monitoring-service/initializers"
	"go-monitoring-service/models"
	"net/http"
)

func Dashboard(c *gin.Context) {
	userInterface, _ := c.Get("user")
	user := userInterface.(models.User)
	var monitors []models.Monitor
	initializers.DB.Where("user_id = ?", user.ID).Find(&monitors)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"user":     user,
		"monitors": monitors,
	})
}
