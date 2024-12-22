package service

import (
	"context"
	"errors"
	"time"

	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
)

type CommentService interface {
	CreateComment(ctx context.Context, photoId uint64, req model.CommentRequest) (model.CommentResponse, error)
	GetAllCommentsByPhotoId(ctx context.Context, photoId uint64) ([]model.Comment, error)
	GetCommentById(ctx context.Context, commentId uint64) (model.Comment, error)
	UpdateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	DeleteComment(ctx context.Context, commmentId uint64) (error)
}

type commentServiceImpl struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentServiceImpl{repo : repo}
}

func (c commentServiceImpl) CreateComment(ctx context.Context, photoId uint64, req model.CommentRequest) (model.CommentResponse, error){
	
	userId, ok := ctx.Value("UserId").(uint32)
	if !ok {
		return model.CommentResponse{}, errors.New("cant convert user Id to uint 64")
	}

	comment := model.Comment{
		Message: req.Message,
		PhotoId: photoId,
		UserId: uint64(userId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := c.repo.CreateComment(ctx, comment)
	if err != nil {
		return model.CommentResponse{}, err
	}

	commentResponse := model.CommentResponse{
		Id: res.Id,
		Message: res.Message,
		PhotoId: res.PhotoId,
		UserId: res.UserId,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return commentResponse, nil
}

func (c commentServiceImpl) GetCommentById(ctx context.Context, commentId uint64) (model.Comment, error){
	comment, err := c.repo.GetCommentById(ctx, commentId)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (c commentServiceImpl) GetAllCommentsByPhotoId(ctx context.Context, photoId uint64) ([]model.Comment, error){
	comments, err := c.repo.GetAllCommentsByPhotoId(ctx, photoId)
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func (c commentServiceImpl) UpdateComment(ctx context.Context, comment model.Comment) (model.Comment, error){
	comment, err := c.repo.UpdateComment(ctx, comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
func (c commentServiceImpl) DeleteComment(ctx context.Context, commmentId uint64) (error){
	err := c.repo.DeleteComment(ctx, commmentId)
	if err != nil {
		return err
	}
	return err
}