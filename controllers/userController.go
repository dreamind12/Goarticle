package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"Gotest/config"
	"Gotest/database"
	"Gotest/models"
)

func RegisterUser(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	userInput.Password = string(hashedPassword)
	userInput.CreatedAt = time.Now()

	// Simpan pengguna ke database
	result := database.DB.Create(&userInput)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "data": userInput})
}

func LoginUser(c *gin.Context) {
	var loginInput models.User
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := database.DB.Where("username = ? OR email = ?", loginInput.Username, loginInput.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	tokenString, err := config.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set cookie with token
	cookie := &http.Cookie{
		Name:     "login_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour), // Expires in 1 day
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{"data": user, "token": tokenString})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserInfoLogin(c *gin.Context) {	
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data pengguna tidak ditemukan"})
		return
	}
	
	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": currentUser})
}

func UpdateUserByID(c *gin.Context) {
	userID := c.Param("id")

	var existingUser models.User
	result := database.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateInput models.User
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateInput.Username != "" {
		existingUser.Username = updateInput.Username
	}

	if updateInput.Email != "" {
		existingUser.Email = updateInput.Email
	}

	if updateInput.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		existingUser.Password = string(hashedPassword)
	}

	if updateInput.Profile != "" {
		// Handle profile image upload
		// You may want to use a library like github.com/gin-gonic/gin#ShouldBindForm for file uploads
		// Save the uploaded file and update the profile field
		// For simplicity, assume the updateInput.Profile is a URL to the uploaded image
		existingUser.Profile = updateInput.Profile
	}

	result = database.DB.Save(&existingUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

func UpdateUserByLogin(c *gin.Context) {	
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data pengguna tidak ditemukan"})
		return
	}
	
	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data pengguna"})
		return
	}

	// Mengambil ID pengguna yang sedang login
	userID := currentUser.ID

	var existingUser models.User
	result := database.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateInput models.User
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateInput.Username != "" {
		existingUser.Username = updateInput.Username
	}

	if updateInput.Email != "" {
		existingUser.Email = updateInput.Email
	}

	if updateInput.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		existingUser.Password = string(hashedPassword)
	}

	if updateInput.Profile != "" {
		existingUser.Profile = updateInput.Profile
	}

	result = database.DB.Save(&existingUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

func LogoutUser(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "login_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}
