package model

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
    Id        uint64         `json:"id"`
    Title     string         `json:"title"`
    Url       string         `json:"url"`
    Caption   string         `json:"caption"`
    UserId    uint64         `json:"user_id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type PhotoRequest struct {
	Title  	  string         `json:"title" validate:"required"`
	Url       string         `json:"url" validate:"required"`
	Caption   string         `json:"caption"`
}

type PhotoResponse struct {
    Id        uint64    `json:"id"`
    Title     string    `json:"title"`
    Url       string    `json:"url"`
    Caption   string    `json:"caption"`
    UserId    uint64    `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    
}