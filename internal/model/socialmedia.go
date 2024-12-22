package model

import (
	"net/url"
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	Id        uint64 		 `json:"id"`
	Name	  string         `json:"name"`
	Url       string 		 `json:"url"`
	UserId    uint64 		 `json:"user_id"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type SocialMediaRequest struct{
	Name	  string         `json:"name" validate:"required,min=3"`
	Url       string 		 `json:"url" validate:"required,min=10"`
}
type SocialMediaResponse struct{
	Id		  uint64		 `json:"id"`
	Name	  string         `json:"name"`
	Url       string 		 `json:"url"`
	UserId    uint64 		 `json:"user_id"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
}

type SocialMediaUpdate struct{
	SocialMediaRequest
	SocialMediaId uint64
	UserId uint64
}

type SocialMediaDelete struct{
	SocialMediaId uint64
	UserId uint64
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func ValidateSocialMediaRequest(req SocialMediaRequest) []string{
	var errorMessages []string
	
	// validate name
	if req.Name == "" {
		errorMessages = append(errorMessages, "name is required")
	} 
	if len(req.Name) <= 3 {
		errorMessages = append(errorMessages, "name should be at least 4 characters long")
	}

	// validate url
	if req.Url == "" {
		errorMessages = append(errorMessages, "url is required")
	} 
	if !isValidURL(req.Url) {
		errorMessages = append(errorMessages, "url must be a valid url")
	}
	return errorMessages
}


