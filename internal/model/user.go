package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64 		 `json:"id"`
	Username  string 		 `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	DoB       time.Time      `json:"dob" gorm:"column:dob"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type UserSignUp struct {
	Username string 	`json:"username" validate:"required"`
	Email    string 	`json:"email" validate:"required"`
	Password string     `json:"password" validate:"required"`
	DoB      time.Time  `json:"dob" validate:"required"` 
}

type UserSignIn struct {
	Email    string 	`json:"email" validate:"required"` 
	Password string     `json:"password" validate:"required"`
}
type UserUpdate struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	DoB      time.Time `json:"dob" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
// func validator
// func (u UserSignUp) Validate() error {
// 	// check username
// 	if u.Username == "" {
// 		return errors.New("invalid username")
// 	}
// 	if len(u.Password) < 6 {
// 		return errors.New("invalid password")
// 	}
// 	return nil
// }
