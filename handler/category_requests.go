package handler

import (
	"errors"
	"fmt"
)

type CreateCategoryRequest struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (r *CreateCategoryRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.Color == "" {
		return errors.New("color is required")
	}

	return nil
}

type UpdateCategoryRequest struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (r *UpdateCategoryRequest) Validate() error {
	if r.Title != "" || r.Color != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
