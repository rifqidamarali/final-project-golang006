package model

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID        uint64 		 `json:"id"`
	Name	  string         `json:"name"`
	Url       string 		 `json:"url"`
	UserID    uint64 		 `json:"user_id"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type SocialMediaCreate struct{
	Name	  string         `json:"name"`
	Url       string 		 `json:"url"`
}
type SocialMediaUpdate struct{
	Name	  string         `json:"name"`
	Url       string 		 `json:"url"`
}



