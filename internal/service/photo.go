package service

import (
	"context"
	"errors"
	"time"

	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, photo model.PhotoRequest) (*model.PhotoResponse, error)
	GetAllPhotosById(ctx context.Context, userId uint64) ([]model.Photo, error)
	GetPhotoById(ctx context.Context, photoId uint64) (model.Photo, error)
	UpdatePhoto(ctx context.Context, photo model.Photo) (model.PhotoResponse, error)
	DeletePhoto(ctx context.Context, photoId uint64) (*model.Photo, error)
}

type photoServiceImpl struct {
	repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{repo: repo}
}

func (p *photoServiceImpl) CreatePhoto(ctx context.Context, req model.PhotoRequest) (*model.PhotoResponse, error) {
	// from uint32 
	
	userId := ctx.Value("UserId")
	uint32Value, ok := userId.(uint32)
    if !ok {
		return nil, errors.New("cant convert to uint32")
    }

	photo := model.Photo{
		Title: req.Title,
		Url: req.Url,
		Caption: req.Caption,
		UserId: uint64(uint32Value),
	}
	
	res, err := p.repo.CreatePhoto(ctx, photo)
	if err != nil {
		return nil, err
	}

	photoResponse := model.PhotoResponse{}
	photoResponse.Id = res.Id
	photoResponse.Title = res.Title
	photoResponse.Caption = res.Caption
	photoResponse.Url = res.Url
	photoResponse.UserId = uint64(uint32Value)

	return &photoResponse, err
}

func (p *photoServiceImpl) GetAllPhotosById(ctx context.Context, userId uint64) ([]model.Photo, error) {
	photos, err := p.repo.GetAllPhotosById(ctx, userId)
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (p *photoServiceImpl) GetPhotoById(ctx context.Context, photoId uint64) (model.Photo, error) {
	photo, err := p.repo.GetPhotoById(ctx, photoId)
	if err != nil {
		return model.Photo{}, err
	}

	return photo, err
}

func (p *photoServiceImpl) UpdatePhoto(ctx context.Context, photo model.Photo) (model.PhotoResponse, error) {
	photoResponse := model.PhotoResponse{}
	photo.UpdatedAt = time.Now()
	photoGet, err := p.repo.GetPhotoById(ctx, photo.Id)
	if err != nil {
		return model.PhotoResponse{}, err
	}
	photo, err = p.repo.UpdatePhoto(ctx, photo)
	if err != nil {
		return photoResponse, err
	}

	photoResponse.Id = photo.Id
	photoResponse.Title = photo.Title
	photoResponse.Caption = photo.Caption
	photoResponse.Url = photo.Url
	photoResponse.UserId = photoGet.UserId
	photoResponse.CreatedAt = photoGet.CreatedAt
	photoResponse.UpdatedAt = photo.UpdatedAt

	return photoResponse, err
}

func (p *photoServiceImpl) DeletePhoto(ctx context.Context, photoId uint64) (*model.Photo, error) {
	photo, err := p.GetPhotoById(ctx, photoId)
	if err != nil {
		return &model.Photo{}, err
	}

	if photo.Id == 0 {
		return &model.Photo{}, err
	}

	err = p.repo.DeletePhoto(ctx, photoId)
	if err != nil {
		return nil, err
	}

	return &photo, nil
}