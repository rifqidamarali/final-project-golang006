package repository

import (
	"context"

	"github.com/rifqidamarali/final-project-golang006/internal/infrastructure"
	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	// GetUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserById(ctx context.Context, id uint64) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUserById(ctx context.Context, id uint64) error
}

// type UserCommand interface {
// 	CreateUser(ctx context.Context, user model.User) (model.User, error)
// }

type userRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewUserRepository(db infrastructure.GormPostgres) UserRepository {
	return &userRepositoryImpl{db: db}
}

// func (u *userQueryImpl) GetUsers(ctx context.Context) ([]model.User, error) {
// 	db := u.db.GetConnection()
// 	users := []model.User{}
// 	if err := db.
// 		WithContext(ctx).
// 		Table("users").
// 		Find(&users).Error; err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
func (u *userRepositoryImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("users").
		Save(&user).Error; err != nil {
			return model.User{}, err
	}
	return user, nil
}

func (u *userRepositoryImpl) GetUserById(ctx context.Context, id uint64) (model.User, error) {
	db := u.db.GetConnection()
	user := model.User{}
	if err := db.
		WithContext(ctx).
		Table("users").
		Where("id = ?", id).
		Find(&user).Error; err != nil {
		// if user not found, return nil error
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}

		return user, err
	}
	return user, nil
}

func (u *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	db := u.db.GetConnection()
	user := model.User{}
	if err := db.
		WithContext(ctx).
		Table("users").
		Where("email = ?", email).
		Find(&user).Error; err != nil {
		// if user not found, return nil error
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (u *userRepositoryImpl) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	db := u.db.GetConnection()

	if err := db.
		WithContext(ctx).
		Updates(&user).
		Error; err != nil {
			return user, err
		}

	return user, nil
}

func (u *userRepositoryImpl) DeleteUserById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("users").
		Delete(&model.User{ID: id}).
		Error; err != nil {
		return err
		}
	return err
}