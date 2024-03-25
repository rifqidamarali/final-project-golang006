package repository

import (
	"context"

	"github.com/rifqidamarali/final-project-golang006/internal/infrastructure"
	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	// GetUsers(ctx context.Context) ([]model.User, error)
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	GetAllPhotosByUserId(ctx context.Context, userId uint64) ([]model.Photo, error)
	GetPhotoById(ctx context.Context, Id uint64) (model.Photo, error)
	UpdatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	DeletePhoto(ctx context.Context, Id uint64) error
}

type photoRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewPhotoRepository(db infrastructure.GormPostgres) PhotoRepository {
	return &photoRepositoryImpl{db: db}
}

func (p *photoRepositoryImpl) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	db := p.db.GetConnection()

	if err := db.
		WithContext(ctx).
		Table("photos").
		Create(&photo).
		Error; err != nil{
			return model.Photo{}, err
		}

	return photo, nil
}

func (p *photoRepositoryImpl) GetAllPhotosByUserId(ctx context.Context, userId uint64) ([]model.Photo, error) {
	db := p.db.GetConnection()
	photos := []model.Photo{}

	err := db.
		WithContext(ctx).
		Table("photos").
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, username").Table("users").Where("deleted_at is null")
		}).
		Find(photos).
		Error

	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (p *photoRepositoryImpl) GetPhotoById(ctx context.Context, id uint64) (model.Photo, error) {
	db := p.db.GetConnection()
	photo := model.Photo{}

	err := db.
		WithContext(ctx).
		Table("photos").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Find(&photo).
		Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (p *photoRepositoryImpl) UpdatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	db := p.db.GetConnection()
	err := db.
		WithContext(ctx).
		Updates(&photo).
		Error

	return photo, err
}

func (p *photoRepositoryImpl) DeletePhoto(ctx context.Context, id uint64) error {
	db := p.db.GetConnection()

	if err := db.
		WithContext(ctx).
		Table("photos").
		Delete(&model.Photo{ID : id}).
		Error; err != nil {
			return err
		}

	return nil
}