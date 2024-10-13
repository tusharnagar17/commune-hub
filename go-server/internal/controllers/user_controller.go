package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tusharnagar17/commune-hub/internal/config"
	"github.com/tusharnagar17/commune-hub/internal/database"
	"github.com/tusharnagar17/commune-hub/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// JWT secret key
var jwtSecret = []byte(config.LoadConfig().JWTSecret)

// CreateJWT generated a new JWT token
func CreateJWT(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Printf("tokenString error %v", err)
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}

func Register(c *gin.Context) {
	log.Println("Register request sented!")
	

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email, and password are required"})
		return
	}

	// Hash the password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error":"Failed to hash password",
		})
		return
	}
	user.Password = hashedPassword

	// saves the user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login request
func Login(c *gin.Context) {
	var loginDetails struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	var user models.User

	if err := database.DB.Where("email = ?", loginDetails.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := CheckPassword(user.Password, loginDetails.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// generate token
	token, err := CreateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token", "full error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})
}