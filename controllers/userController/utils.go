package userController

import (
	"github.com/gin-gonic/gin"
	"go-monitoring-service/initializers"
	"go-monitoring-service/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashed), err
}

func UsersShow(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg":    "Show all users",
		"users":  users,
	})
}

func UserShow(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	initializers.DB.First(&user, id)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg":    "Show specific user",
		"user":   user,
	})
}

func UserUpdate(c *gin.Context) {
	var userBase struct {
		Name     string
		Email    string
		Password string
	}
	err := c.Bind(&userBase)
	if err != nil {
		return
	}

	var user models.User
	id := c.Param("id")
	initializers.DB.First(&user, id)
	initializers.DB.Model(&user).Updates(models.User{Name: userBase.Name, Password: userBase.Password, Email: userBase.Email})
	c.JSON(200, gin.H{
		"status": "ok",
		"msg":    "user updated successfully",
		"user":   user,
	})
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.User{}, id)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg":    "user deleted successfully",
	})
}
