package size

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
)

func GetSizesHandler(c *gin.Context) {
	var sizes []models.Size
	if err := db.Db.Find(&sizes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sizes)
}

func GetSizeHandler(c *gin.Context) {
	var size models.Size
	if err := db.Db.First(&size, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, size)
}

func CreateSizeHandler(c *gin.Context) {
	var size models.Size
	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Create(&size).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, size)
}

func UpdateSizeHandler(c *gin.Context) {
	var size models.Size
	if err := db.Db.First(&size, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Save(&size).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, size)
}

func DeleteSizeHandler(c *gin.Context) {
	var size models.Size
	if err := db.Db.First(&size, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Delete(&size).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "size deleted successfully"})
}
