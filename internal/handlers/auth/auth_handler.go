package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"sneakers-app/internal/models"
	"sneakers-app/pkg/db"
	"time"
)

func RegisterHandler(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := isUsernameUnique(body.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := isEmailUnique(body.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("_____________", body)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to hash password"})
		return
	}

	body.Password = string(hashedPassword)

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		Role:     "client",
	}

	if err := db.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func isUsernameUnique(username string) error {
	var existingUser models.User
	if err := db.Db.First(&existingUser, "username = ?", username).Error; err == nil {
		return err
	}
	return nil
}

func isEmailUnique(email string) error {
	var existingUser models.User
	if err := db.Db.First(&existingUser, "email = ?", email).Error; err == nil {
		return err
	}
	return nil
}

func LoginHandler(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := db.Db.First(&existingUser, "username = ?", body.Username).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password", "details": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password", "details": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  existingUser.ID,
		"name": existingUser.Username,
		"role": existingUser.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"role": existingUser.Role, "token": tokenString, "message": "user logged in successfully"})
}
