package service

import (
	"context"

	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
)

type SocialMediaService interface {
	CreateSocialMedia(ctx context.Context, userId uint64, SocialMediaCreate model.SocialMedia) (model.SocialMedia, error)
	GetAllSocialMediasByUserId(ctx context.Context, userId uint64) ([]model.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, socialMediaId uint64) (model.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, socialMediaId model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, socialMediaId uint64) error
}

type socialMediaServiceImpl struct {
	repo repository.SocialMediaRepository
}

func NewSocialMediaService(repo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{repo: repo}
}

func (s *socialMediaServiceImpl) CreateSocialMedia(ctx context.Context, SocialMediaCreate model.SocialMediaCreate) (model.SocialMedia, error) {
	socialMedia := model.SocialMedia{}
	socialMedia.Name = SocialMediaCreate.Name
	socialMedia.Url = SocialMediaCreate.Url
	// socialMedia.UserID = userId

	res, err := s.repo.CreateSocialMedia(ctx, socialMedia)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s *socialMediaServiceImpl) GetAllSocialMediasByUserId(ctx context.Context, userId uint64) ([]model.SocialMedia, error) {
	socialMedia, err := s.repo.GetAllSocialMediasByUserId(ctx, userId)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, err
}

func (s *socialMediaServiceImpl) GetSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error) {
	socialMedia, err := s.repo.GetSocialMediaById(ctx, id)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (s *socialMediaServiceImpl) UpdateSocial(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	socialMedia, err := s.repo.UpdateSocialMedia(ctx, socialMedia)

	if err != nil {
		return socialMedia, err
	}

	res := model.SocialMedia{}
	res.ID = socialMedia.ID
	res.UserID = socialMedia.UserID
	res.Name = socialMedia.Name
	res.Url = socialMedia.Url
	res.UpdatedAt = socialMedia.UpdatedAt

	return res, nil
}

func (s *socialMediaServiceImpl) DeleteSocial(ctx context.Context, id uint64) error {
	err := s.repo.DeleteSocialMedia(ctx, id)
	return err
}