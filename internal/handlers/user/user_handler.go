package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"path/filepath"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
)

func CreateUser(c *gin.Context) {
	var user models.User

	file, err := c.FormFile("photo")
	if err != nil {
		if user.UserPhoto != nil {
			os.Remove("../media/" + *user.UserPhoto)
		}

		var photoFIleName string
		if file != nil {
			uuidFileName := uuid.New().String()
			extension := filepath.Ext(file.Filename)
			photoFIleName = uuidFileName + extension

			if err := c.SaveUploadedFile(file, "../media/"+photoFIleName); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			user.UserPhoto = &photoFIleName
		}
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserById(c *gin.Context) {
	var user models.User
	if err := db.Db.Unscoped().Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := db.Db.Unscoped().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := db.Db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		if user.UserPhoto != nil {
			os.Remove("../media/" + *user.UserPhoto)
		}

		var photoFIleName string
		if file != nil {
			uuidFileName := uuid.New().String()
			extension := filepath.Ext(file.Filename)
			photoFIleName = uuidFileName + extension

			if err := c.SaveUploadedFile(file, "../media/"+photoFIleName); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			user.UserPhoto = &photoFIleName
		}
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.Db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := db.Db.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.UserPhoto != nil {
		os.Remove("../media/" + *user.UserPhoto)
	}

	if err := db.Db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func RestoreUser(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	// Проверяем, существует ли пользователь с указанным ID
	if err := db.Db.Unscoped().First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Проверяем, был ли пользователь удален (поле deleted_at не пустое)
	if !user.DeletedAt.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not deleted"})
		return
	}

	// Если пользователь был удален, сбрасываем поле deleted_at
	if err := db.Db.Model(&user).Update("deleted_at", nil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User restored successfully"})
}

func GetMaxUserId(c *gin.Context) {
	var maxId int

	if err := db.Db.Model(&models.User{}).Select("MAX(id)").Scan(&maxId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"maxId": maxId})
}
