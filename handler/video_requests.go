package handler

import (
	"errors"
	"fmt"
)

type CreateVideoRequest struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	URL string `json:"url" binding:"required"`
	CategoryID  uint `json:"category_id"`
}

func (r *CreateVideoRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	} 

	if r.Description == "" {
		return errors.New("description is required")
	}

	if r.URL == "" {
		return errors.New("url is required")
	}

	return nil
}

type UpdateVideoRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	URL string `json:"url"`
}

func (r *UpdateVideoRequest) Validate() error {
	if r.Title != "" || r.Description != "" || r.URL != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

