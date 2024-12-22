package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
)

type SocialMediaServcie interface {
	CreateSocialMedia(ctx context.Context, req model.SocialMediaRequest, userId uint64) (model.SocialMediaResponse, error )
	GetAllSocialMediasByUserId(ctx context.Context, userId uint64) ([]model.SocialMediaResponse, error)
	GetSocialMediaById(ctx context.Context, socialId uint64) (model.SocialMediaResponse, error)
	UpdateSocialMedia(ctx context.Context, input model.SocialMediaUpdate) (model.SocialMediaResponse, error)
	DeleteSocialMedia(ctx context.Context, input model.SocialMediaDelete) error
}

type socialMediaServiceImpl struct {
	repository repository.SocialMediaRepository
}

func NewSocialMediaService(repository repository.SocialMediaRepository) SocialMediaServcie {
	return &socialMediaServiceImpl{repository: repository}
}

func(s *socialMediaServiceImpl) CreateSocialMedia(ctx context.Context, req model.SocialMediaRequest, userId uint64) (model.SocialMediaResponse, error){
	socialMedia := model.SocialMedia{
		Name: req.Name,
		Url: req.Url,
		UserId: userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := s.repository.CreateSocialMedia(ctx, socialMedia)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	socialMediaResponse := model.SocialMediaResponse{
		Id: res.Id,
		Name: res.Name,
		Url: res.Url,
		UserId: res.UserId,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return socialMediaResponse, nil
}

func (s *socialMediaServiceImpl) GetAllSocialMediasByUserId(ctx context.Context, userId uint64) ([]model.SocialMediaResponse, error){
	socialMedias, err := s.repository.GetAllSocialMediasByUserId(ctx, userId)
	if err != nil {
		return []model.SocialMediaResponse{}, err
	}
	//append use when the len of data unpredicted?
	// res := make([]model.SocialMediaResponse, 0, len(socialMedias))
	// for _, sm := range socialMedias {
    //     res= append(res, model.SocialMediaResponse{
    //         Id:   sm.Id,
    //         Name: sm.Name,
	// 		CreatedAt: sm.CreatedAt,
	// 		UpdatedAt: sm.UpdatedAt,
    //     })
    // }
	// use it when len of data are predicted
	res := make([]model.SocialMediaResponse, len(socialMedias))
	for i, sm := range socialMedias {
		res[i] = model.SocialMediaResponse{
			Id: sm.Id,
			Name: sm.Name,
			Url: sm.Url,
			UserId: sm.UserId,
			CreatedAt: sm.CreatedAt,
			UpdatedAt: sm.UpdatedAt,
		}
	}

	return res, nil
}
// 
func (s *socialMediaServiceImpl) GetSocialMediaById(ctx context.Context, socialMediaId uint64) (model.SocialMediaResponse, error){
	socialMedia, err := s.repository.GetSocialMediaById(ctx, socialMediaId)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}
	res := model.SocialMediaResponse{
		Id: socialMedia.Id,
		Name: socialMedia.Name,
		Url: socialMedia.Name,
		UserId: socialMedia.UserId,
		CreatedAt: socialMedia.CreatedAt,
		UpdatedAt: socialMedia.UpdatedAt,

	}

	return res, err
}

func (s *socialMediaServiceImpl) UpdateSocialMedia(ctx context.Context, input model.SocialMediaUpdate) (model.SocialMediaResponse, error){
	socialMediaGet, err := s.repository.GetSocialMediaById(ctx, input.SocialMediaId)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}
	// If social media has been deleted --> model.id == 0
    if socialMediaGet.Id == 0 {
        return model.SocialMediaResponse{}, NewCustomError(ErrNotFound, fmt.Sprintf("social media with ID %d does not exist", input.SocialMediaId))
    }

    // If user ID doesn't match, unauthorized
    if socialMediaGet.UserId != input.UserId {
        return model.SocialMediaResponse{}, NewCustomError(ErrForbidden, fmt.Sprintf("user %d cannot edit social media owned by user %d", input.UserId, socialMediaGet.UserId))
    }
	socialMedia := model.SocialMedia{
		Id: socialMediaGet.Id,
		Name: input.Name,
		Url: input.Url,
		UserId: socialMediaGet.UserId,
		UpdatedAt: time.Now(),

	}

	socialMedia, err = s.repository.UpdateSocialMedia(ctx, socialMedia)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	res := model.SocialMediaResponse{
		Id: socialMedia.Id,
		Name: socialMedia.Name,
		Url: socialMedia.Url,
		UserId: socialMedia.UserId,
		CreatedAt: socialMediaGet.CreatedAt,
		UpdatedAt: socialMedia.UpdatedAt,
	}

	return res, err
}

func (s *socialMediaServiceImpl) DeleteSocialMedia(ctx context.Context, input model.SocialMediaDelete) error {
    // Get by ID
    socialMediaGet, err := s.repository.GetSocialMediaById(ctx, input.SocialMediaId)
    if err != nil {
        return err
    }

    // fmt.Printf("socialMediaGet: %+v\n", socialMediaGet)
    // fmt.Printf("input.UserId: %d\n", input.UserId)

    // If social media has been deleted --> model.id == 0
    if socialMediaGet.Id == 0 {
        return NewCustomError(ErrNotFound, fmt.Sprintf("social media with ID %d does not exist", input.SocialMediaId))
    }

    // If user ID doesn't match, unauthorized
    if socialMediaGet.UserId != input.UserId {
        return NewCustomError(ErrForbidden, fmt.Sprintf("user %d cannot delete social media owned by user %d", input.UserId, socialMediaGet.UserId))
    }

    if err := s.repository.DeleteSocialMedia(ctx, input.SocialMediaId); err != nil {
        return err
    }

    return nil
}
