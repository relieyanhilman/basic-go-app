package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/relieyanhilman/basic-go-app/models"

	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

// Create Post Handler
func (pc *PostController) CreatePost(ctx *gin.Context) {
	//Take the param
	currentUser := ctx.MustGet("currentUser").(models.User)

	//Error check for validating the request.body
	var payload *models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//Setup new post
	now := time.Now()
	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	//Create new post and error checking
	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	//return to the client the new post
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

// Update Post Handler
func (pc *PostController) UpdatePost(ctx *gin.Context) {

	//Take the params
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	//error check for the request.body
	var payload *models.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	//get the past post and error check for the past post
	var updatedPost models.Post
	result := pc.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	//setup new post to be updated
	now := time.Now()
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: now,
	}

	//update new post to the past post
	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	//return to the client the new updated post
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})

}

// Get single post handler
func (pc *PostController) FindPostById(ctx *gin.Context) {

	//take the param
	postId := ctx.Param("postId")

	//get single post and error checking
	var post models.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	//return to the client the post
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

// get all post handler
func (pc *PostController) FindPosts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	//convert string to int with strconv
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)

	//define how many records that will be skipped before starting to return the records
	offset := (intPage - 1) * intLimit

	//get all post and error checking
	var posts []models.Post
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	//return to the client status, results, and post data
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

//delete post handler
func (pc *PostController) DeletePost(ctx *gin.Context){
	//take the param
	postId := ctx.Param("postId")

	//delete post
	result := pc.DB.Delete(&models.Post{}, "id = ?", postId)

	//error checking
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

	
}
