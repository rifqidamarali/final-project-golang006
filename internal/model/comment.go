package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id        uint64 		 `json:"id"`
	Message   string 		 `json:"message"`
	PhotoId   uint64 		 `json:"photo_id"`
	UserId    uint64 		 `json:"user_id"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommentResponse struct {
	Id        uint64 		 `json:"id"`
	Message   string 		 `json:"message" validate:"required"`
	PhotoId   uint64 		 `json:"photo_id"`
	UserId    uint64 		 `json:"user_id"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
}