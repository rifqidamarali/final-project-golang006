package service

import (
	"context"
	"time"

	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	GetAllPhotosByUserId(ctx context.Context, userId uint64) ([]model.Photo, error)
	GetPhotoById(ctx context.Context, photoId uint64) (model.Photo, error)
	UpdatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	DeletePhoto(ctx context.Context, photoId uint64) error
}

type photoServiceImpl struct {
	repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{repo: repo}
}

func (p *photoServiceImpl) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	photo, err := p.repo.CreatePhoto(ctx, photo)
	if err != nil {
		return photo, err
	}

	photoResponse := model.Photo{}
	photoResponse.ID = photo.ID
	photoResponse.Title = photo.Title
	photoResponse.Caption = photo.Caption
	photoResponse.Url = photo.Url
	photoResponse.UserID = photo.UserID
	photoResponse.CreatedAt = time.Now()

	return photoResponse, err
}

func (p *photoServiceImpl) GetAllPhotosByUserId(ctx context.Context, userId uint64) ([]model.Photo, error) {
	photos, err := p.repo.GetAllPhotosByUserId(ctx, userId)
	if err != nil {
		return photos, err
	}

	return photos, err
}

func (p *photoServiceImpl) GetPhotoById(ctx context.Context, photoId uint64) (model.Photo, error) {
	photo, err := p.repo.GetPhotoById(ctx, photoId)
	if err != nil {
		return model.Photo{}, err
	}

	return photo, err
}

func (p *photoServiceImpl) UpdatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	photo, err := p.repo.UpdatePhoto(ctx, photo)
	if err != nil {
		return photo, err
	}

	photoResponse := model.Photo{}
	photoResponse.ID = photo.ID
	photoResponse.Title = photo.Title
	photoResponse.Caption = photo.Caption
	photoResponse.Url = photo.Url
	photoResponse.UserID = photo.UserID
	photoResponse.UpdatedAt = photo.UpdatedAt

	return photoResponse, err}

func (p *photoServiceImpl) DeletePhoto(ctx context.Context, photoId uint64) error {
	err := p.repo.DeletePhoto(ctx, photoId)

	return err
}