package sneakers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
)

const CloudinaryUrl = "cloudinary://497718144425551:SH9dXr5R5FG6I7w_IlAuzaAXXLs@dlvk8igke"

func GetSneakersHandler(c *gin.Context) {
	var sneakers []models.Sneaker
	if err := db.Db.Unscoped().Find(&sneakers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sneakers", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneakers)
}

func GetSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker
	if err := db.Db.First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sneaker", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func CreateSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker

	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON", "details": err.Error()})
		return
	}

	//file, _, err := c.Request.FormFile("photo")
	//if err != nil {
	//	log.Println(file)
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Photo is required", "details": err.Error()})
	//	return
	//}
	//
	//ctx := context.Background()
	//
	//cldServic, _ := cloudinary.NewFromURL(CloudinaryUrl)
	//resp, _ := cldServic.Upload.Upload(ctx, file, uploader.UploadParams{})
	//log.Println(resp.SecureURL)
	//
	//sneaker.Photo = resp.SecureURL

	if err := db.Db.Create(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sneaker", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sneaker)
}

func UpdateSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker
	if err := db.Db.First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find sneaker for update", "details": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON for update", "details": err.Error()})
		return
	}
	if err := db.Db.Save(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sneaker", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sneaker)
}

func DeleteSneakerHandler(c *gin.Context) {
	var sneaker models.Sneaker
	if err := db.Db.First(&sneaker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find sneaker for delete", "details": err.Error()})
		return
	}
	if err := db.Db.Delete(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sneaker", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sneaker deleted successfully"})
}
