package purchaseHistory

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
)

func GetPurchaseHistoryHandler(c *gin.Context) {
	var purchaseHistory []models.PurchaseHistory
	if err := db.Db.Find(&purchaseHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchaseHistory)
}

func GetPurchaseHistoryByUserHandler(c *gin.Context) {
	var purchaseHistory []models.PurchaseHistory
	if err := db.Db.Where("user_id = ?", c.Param("id")).Find(&purchaseHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchaseHistory)
}

func GetPurchaseHistoryByProductHandler(c *gin.Context) {
	var purchaseHistory []models.PurchaseHistory
	if err := db.Db.Where("product_id = ?", c.Param("id")).Find(&purchaseHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchaseHistory)
}

func CreatePurchaseHistoryHandler(c *gin.Context) {
	var purchaseHistory models.PurchaseHistory
	if err := c.ShouldBindJSON(&purchaseHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Create(&purchaseHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchaseHistory)
}

func UpdatePurchaseHistoryHandler(c *gin.Context) {
	var purchaseHistory models.PurchaseHistory
	if err := db.Db.First(&purchaseHistory, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&purchaseHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Db.Save(&purchaseHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchaseHistory)
}

func DeletePurchaseHistoryHandler(c *gin.Context) {
	var purchaseHistory models.PurchaseHistory
	if err := db.Db.First(&purchaseHistory, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Db.Delete(&purchaseHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchaseHistory)
}
