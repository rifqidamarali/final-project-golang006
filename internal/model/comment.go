package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint64 		 `json:"id"`
	Message   string 		 `json:"message"`
	PhotoID   uint64 		 `json:"photo_id"`
	UserID    uint64 		 `json:"user_id"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type CommentCreate struct {
	PhotoId uint64 `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type CommentUpdate struct {
	PhotoId uint64 `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}