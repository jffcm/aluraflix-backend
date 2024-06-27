package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/aluraflix-backend/schemas"
)

func CreateCategoryHandler(ctx *gin.Context) {
	var request CreateCategoryRequest

	if err := ctx.BindJSON(&request); err != nil {
		log.Printf("Error decoding JSON request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Failed to process request data")
		return
	}

	if err := request.Validate(); err != nil {
		log.Printf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	category := schemas.Category{
		Title: request.Title,
		Color: request.Color,
	}

	if err := db.Create(&category).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error creating category in database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-category", category)
}

func ListCategoriesHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5
	offset := (page - 1) * limit

	var categories []schemas.Category

	if err := db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing categories")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-categories", categories)
}

func ShowCategoryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Invalid ID: ID parameter is missing or empty")
		return
	}

	var category schemas.Category
	if err := db.First(&category, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("category with id: %s not found", id))
		return
	}
	sendSuccess(ctx, http.StatusOK, "show-category", category)
}

func ListVideosByCategoryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid ID: ID parameter is not a valid integer")
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	
	limit := 5
	offset := (page - 1) * limit

	var videos []schemas.Video
	if err := db.Limit(limit).Offset(offset).Where("category_id = ?", categoryID).Find(&videos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing videos")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-videos", videos)
}

func UpdateCategoryHandler(ctx *gin.Context) {
	var request UpdateCategoryRequest

	if err := ctx.BindJSON(&request); err != nil {
		log.Printf("Error decoding JSON request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Failed to process request data")
		return
	}

	if err := request.Validate(); err != nil {
		log.Printf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Query parameter 'id' is required")
		return
	}

	var category schemas.Category
	if err := db.First(&category, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("category with id: %s not found", id))
		return
	}

	if request.Title != "" {
		category.Title = request.Title
	}

	if request.Color != "" {
		category.Color = request.Color
	}

	if err := db.Save(&category).Error; err != nil {
		log.Printf("error updating category: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating category")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-category", category)
}

func DeleteCategoryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Invalid ID: ID parameter is missing or empty")
		return
	}

	var category schemas.Category
	if err := db.First(&category, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("category with id: %s not found", id))
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting category with id: %s", id))
		return
	}

	ctx.Status(http.StatusNoContent)
}
