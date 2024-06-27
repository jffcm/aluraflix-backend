package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/aluraflix-backend/schemas"
)

func CreateVideoHandler(ctx *gin.Context) {
	var request CreateVideoRequest

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

	video := schemas.Video{
		Title: request.Title,
		Description: request.Description,
		URL: request.URL,
		CategoryID: request.CategoryID,
	}

	if err := db.Create(&video).Error; err != nil {
		log.Printf("error creating video: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating video in database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-video", video)
}

func ListVideosHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5
	offset := (page - 1) * limit

	var videos []schemas.Video
	if err := db.Limit(limit).Offset(offset).Preload("Category").Find(&videos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing videos")
		return
	}
	sendSuccess(ctx, http.StatusOK, "list-videos", videos)
}

func FreeVideosHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5
	offset := (page - 1) * limit

	var videos []schemas.Video
	if err := db.Limit(limit).Offset(offset).Preload("Category").Find(&videos).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing videos")
		return
	}
	sendSuccess(ctx, http.StatusOK, "list-videos", videos)
}

func ShowVideoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Invalid ID: ID parameter is missing or empty")
		return
	}

	var video schemas.Video
	if err := db.First(&video, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("video with id: %s not found", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "show-video", video)
}

func UpdateVideoHandler(ctx *gin.Context) {
	var request UpdateVideoRequest

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

	var video schemas.Video
	if err := db.First(&video, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("video with id: %s not found", id))
		return
	}

	if request.Title != "" {
		video.Title = request.Title
	}

	if request.Description != "" {
		video.Description = request.Description
	}

	if request.URL != "" {
		video.URL = request.URL
	}

	if err := db.Save(&video).Error; err != nil {
		log.Printf("error updating video: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating video")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-video", video)
}

func DeleteVideoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Invalid ID: ID parameter is missing or empty")
		return
	}

	var video schemas.Video
	if err := db.First(&video, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("video with id: %s not found", id))
		return
	}

	if err := db.Delete(&video).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting video with id: %s", id))
		return
	}

	ctx.Status(http.StatusNoContent)
}