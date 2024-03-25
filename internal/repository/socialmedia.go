package repository

import (
	"context"

	"github.com/rifqidamarali/final-project-golang006/internal/infrastructure"
	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(ctx context.Context, social model.SocialMedia) (model.SocialMedia, error )
	GetAllSocialMediasByUserId(ctx context.Context, userId uint64) ([]model.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, socialId uint64) (model.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, social model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, socialId uint64) error
}

type socialMediaRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewSocialMediaRepository(db infrastructure.GormPostgres) SocialMediaRepository {
	return &socialMediaRepositoryImpl{db: db}
}

func (s *socialMediaRepositoryImpl) CreateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	db := s.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("social_medias").
		Create(&socialMedia).Error; err!= nil {
			return model.SocialMedia{}, err
		}
	return socialMedia, nil
}

func (s *socialMediaRepositoryImpl) GetAllSocialMediasByUserId(ctx context.Context, userId uint64) ([]model.SocialMedia, error) {
	db := s.db.GetConnection()
	socialMedias := []model.SocialMedia{}

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, username").Table("users").Where("deleted_at is null")
		}).
		Find(&socialMedias).
		Error

	if err != nil {
		return socialMedias, err
	}

	return socialMedias, nil
}

func (s *socialMediaRepositoryImpl) GetSocialMediaById(ctx context.Context, socialMediaId uint64) (model.SocialMedia, error) {
	db := s.db.GetConnection()
	socialMedia := model.SocialMedia{}

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Where("id = ?", socialMediaId).
		Where("deleted_at IS NULL").
		Find(socialMedia).
		Error

	if err != nil {
		return socialMedia, err
	}

	return socialMedia, err
}

func (s *socialMediaRepositoryImpl) UpdateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	db := s.db.GetConnection()
	err := db.
		WithContext(ctx).
		Table("social_medias").
		Updates(&socialMedia).
		Error

	return socialMedia, err
}

func (s *socialMediaRepositoryImpl) DeleteSocialMedia(ctx context.Context, socialMediaId uint64) error {
	db := s.db.GetConnection()
	social := model.SocialMedia{ID: socialMediaId}

	if err := db.
		WithContext(ctx).
		Table("social_medias").
		Delete(&social).
		Error; err != nil {
			return err
		}

	return nil
}