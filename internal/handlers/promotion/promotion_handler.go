package promotion

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
	"time"
)

func GetPromotionsHandler(c *gin.Context) {
	var promotions []models.Promotion
	if err := db.Db.Unscoped().Find(&promotions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promotions)

}

func GetPromotionHandler(c *gin.Context) {
	var promotion models.Promotion
	if err := db.Db.First(&promotion, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promotion)
}

func CreatePromotionHandler(c *gin.Context) {
	var promotion models.Promotion
	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedDate := time.Now().Format("02.01.2006 15:04:05")
	promotion.ExpirationDate = parsedDate

	if err := db.Db.Create(&promotion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promotion)
}

func UpdatePromotionHandler(c *gin.Context) {
	var promotion models.Promotion
	if err := db.Db.First(&promotion, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Db.Save(&promotion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promotion)
}

func DeletePromotionHandler(c *gin.Context) {
	var promotion models.Promotion
	if err := db.Db.First(&promotion, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Db.Delete(&promotion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promotion)
}
