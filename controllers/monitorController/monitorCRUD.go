package monitorController

import (
	"github.com/gin-gonic/gin"
	"go-monitoring-service/initializers"
	"go-monitoring-service/models"
	"net/http"
	"strconv"
	"time"
)

func AddMonitor(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	user := userInterface.(models.User)

	name := c.PostForm("Name")
	monitorType := c.PostForm("Type")
	target := c.PostForm("Target")
	method := c.PostForm("Method")
	expectedStatusStr := c.PostForm("ExpectedStatus")
	keyword := c.PostForm("Keyword")
	intervalStr := c.PostForm("Interval")
	retriesStr := c.PostForm("Retries")
	config := c.PostForm("Config")
	activeStr := c.PostForm("Active")

	expectedStatus, _ := strconv.Atoi(expectedStatusStr)
	intervalSec, _ := strconv.Atoi(intervalStr)
	retries, _ := strconv.Atoi(retriesStr)

	active := activeStr == "on"

	monitor := models.Monitor{
		UserID:         user.ID,
		Name:           name,
		Type:           monitorType,
		Target:         target,
		Method:         method,
		ExpectedStatus: expectedStatus,
		Keyword:        keyword,
		Interval:       time.Duration(intervalSec) * time.Second,
		Retries:        retries,
		Config:         config,
		Active:         active,
	}

	if err := initializers.DB.Create(&monitor).Error; err != nil {
		c.HTML(http.StatusBadRequest, "add-monitor.html", gin.H{
			"Error": "Failed to create monitor",
		})
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

func UpdateMonitor(c *gin.Context) {}

func DeleteMonitor(c *gin.Context) {}
