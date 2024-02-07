package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jicodes/go-crud/initializers"
	"github.com/jicodes/go-crud/models"
)

func CreatePost(c *gin.Context) {
	// get data from request body and bind it to a struct
	var body struct {
		Title string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Content: body.Content}
	result := initializers.DB.Create(&post)
	
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post" : post,
	})
	
}

func GetPosts(c *gin.Context) {
	// get all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return the posts
	c.JSON(200, gin.H{
		"posts" : posts,
	})
}

func GetPost(c *gin.Context) {
	// get a post
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post" : post,
	})
}

func UpdatePost(c *gin.Context) {
	// get data from request body and bind it to a struct
	var body struct {
		Title string `json:"title"`
		Content string `json:"content"`
	}

	c.Bind(&body)

	// find the post we want to update
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	// update that post

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Content: body.Content,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post" : post,
	})
}

func DeletePost(c *gin.Context) {
	// delete that post using the id from the request
	var post models.Post
	initializers.DB.Delete(&post, c.Param("id"))

	// return the post
	c.JSON(200, gin.H{
		"deleted: post" : post,
	})
}