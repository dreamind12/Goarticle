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
		// Validasi gagal
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi menggunakan validator
	if err := validate.Struct(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi tambahan untuk status
	if post.Status != "Publish" && post.Status != "Draft" && post.Status != "Trash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post status"})
		return
	}

	// Validasi tambahan untuk title, content, dan category
	if len(post.Title) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title must be at least 20 characters"})
		return
	}

	if len(post.Descript) < 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Descript must be at least 200 characters"})
		return
	}

	if len(post.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category must be at least 3 characters"})
		return
	}

	// Buat post di database
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": post})
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

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	// Ambil semua data dari tabel posts
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
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

	// Validasi tambahan untuk title, content, dan category
	if len(newPost.Title) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title must be at least 20 characters"})
		return
	}

	if len(newPost.Descript) < 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Descript must be at least 200 characters"})
		return
	}

	if len(newPost.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category must be at least 3 characters"})
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

	c.JSON(http.StatusOK, gin.H{"data": existingPost})
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
