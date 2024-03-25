package model

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint64         `json:"id"`
	Title  	  string         `json:"title"`
	Url       string         `json:"url"`
	Caption   string         `json:"caption"`
	UserID    uint64 		 `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}