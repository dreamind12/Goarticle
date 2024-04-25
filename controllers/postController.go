package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"Gotest/database"
	"Gotest/models"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}
	userData := user.(models.User)
	post.User_ID = userData.ID
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func GetPostByID(c *gin.Context) {
	postID := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching post"})
		}
		return
	}
	userID := post.User_ID
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post, "user": user})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post	
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}

	var responseData []gin.H
	for _, post := range posts {
		userID := post.User_ID
		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
			}
			return
		}

		postData := gin.H{
			"ID":            post.ID,
			"User_ID":       post.User_ID,
			"UserUpdate_ID": post.UserUpdate_ID,
			"Title":         post.Title,
			"Descript":      post.Descript,
			"Category":      post.Category,
			"Status":        post.Status,
			"CreatedAt":     post.CreatedAt,
			"UpdatedAt":     post.UpdatedAt,
			"users":         []models.User{user}, // Include user data under each post
		}

		responseData = append(responseData, postData)
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

func GetPostsWithPaging(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	// Konversi limit dan offset ke integer
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	var posts []models.Post
	if err := database.DB.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func UpdatePostByID(c *gin.Context) {
	postID := c.Param("id")

	var newPost models.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi menggunakan validator
	if err := validate.Struct(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi tambahan untuk status
	if newPost.Status != "Publish" && newPost.Status != "Draft" && newPost.Status != "Trash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post status"})
		return
	}	

	// Update post di database
	var existingPost models.Post
	if err := database.DB.First(&existingPost, postID).Error; err != nil {
		return
	}

	if err := database.DB.Model(&existingPost).Updates(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message" : "Post updated successfully","data": existingPost})
}

func DeletePostByID(c *gin.Context) {
	postID := c.Param("id")

	if err := database.DB.Delete(&models.Post{}, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
