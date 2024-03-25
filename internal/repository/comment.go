package repository

import (
	"context"

	"github.com/rifqidamarali/final-project-golang006/internal/infrastructure"
	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	GetAllCommentsByPhotoId(ctx context.Context, photoId uint64) ([]model.Comment, error)
	GetCommentById(ctx context.Context, commentId uint64) (model.Comment, error)
	UpdateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	DeleteComment(ctx context.Context, commentId uint64) error
}

type commentRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewCommentRepository(db infrastructure.GormPostgres) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (c *commentRepositoryImpl) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	db := c.db.GetConnection()

	if err := db.
		WithContext(ctx).
		Table("comments").
		Create(&comment).
		Error; err != nil {
			return model.Comment{}, err
	}

	return comment, nil
}

func (c *commentRepositoryImpl) GetAllCommentsByPhotoId(ctx context.Context, photoId uint64) ([]model.Comment, error) {
	db := c.db.GetConnection()
	comments := []model.Comment{}

	err := db.
		WithContext(ctx).
		Table("comments").
		Where("photo_id = ?", photoId).
		Where("deleted_at IS NULL").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, username").Table("users").Where("deleted_at is null")
		}).
		Preload("Photo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, title, caption, photo_url, user_id").Table("photos").Where("deleted_at is null")
		}).
		Find(&comments).
		Error

	if err != nil {
		return comments, err
	}

	return comments, err
}

func (c *commentRepositoryImpl) GetCommentById(ctx context.Context, id uint64) (model.Comment, error) {
	db := c.db.GetConnection()
	comment := model.Comment{}

	err := db.
		WithContext(ctx).
		Table("comments").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Find(comment).
		Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (c *commentRepositoryImpl) UpdateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	db := c.db.GetConnection()
	err := db.
		WithContext(ctx).
		Updates(&comment).
		Error

	return comment, err
}

func (c *commentRepositoryImpl) DeleteComment(ctx context.Context, commentId uint64) error {
	db := c.db.GetConnection()
	comment := model.Comment{ID: commentId}

	if err := db.
		WithContext(ctx).
		Model(comment).
		Delete(comment).
		Error; err != nil {
			return err
		}
	return nil
}