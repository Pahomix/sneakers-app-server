package sneakers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
)

func GetSneakersHandler(c *gin.Context) {
	var sneakers []models.Sneaker
	if err := db.Db.Find(&sneakers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneakers)
}

func GetSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker
	if err := db.Db.First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func CreateSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Create(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func UpdateSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker
	if err := db.Db.First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Save(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}
