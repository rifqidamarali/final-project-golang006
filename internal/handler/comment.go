package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/service"
	"github.com/rifqidamarali/final-project-golang006/pkg"
)

type CommentHandler interface {
	CreateComment(ctx *gin.Context)
	GetAllCommentsByPhotoId(ctx *gin.Context)
	GetCommentById(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)


}

type commentHandlerImpl struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) CommentHandler {
	return &commentHandlerImpl{
		service: service,
	}
}

func (c *commentHandlerImpl) CreateComment(ctx *gin.Context){
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid request param"})
		return
	}
	req := model.CommentRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "failed binding json to struct"})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	userId, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "nnt diganti"})
		return
	}
	stdCtx := context.WithValue(ctx.Request.Context(), "UserId", userId)
	res, err := c.service.CreateComment(stdCtx, uint64(photoId), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, pkg.APIResponse{
		Status: "Success",
		Message: "Created",
		Data: res,

	})

}

func (c commentHandlerImpl) GetAllCommentsByPhotoId(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId")) 	
	if err != nil || photoId == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comments, err := c.service.GetAllCommentsByPhotoId(ctx, uint64(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return		
	}	
	
	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "success",
		Message: "Ok",
		Data: comments,
		
	})
}

func (c commentHandlerImpl) GetCommentById(ctx *gin.Context){
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil || commentId == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.APIResponse{
			Status: "failed",
			Message: "",
			Errors: "invalid param",
		})
		return
	}

	comment, err := c.service.GetCommentById(ctx, uint64(commentId))
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, pkg.APIResponse{
			Status: "failed",
			Message: "",
			Errors: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "success",
		Message: "Ok",
		Data: comment,
	})
}

func (c commentHandlerImpl)UpdateComment(ctx *gin.Context){
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil || commentId == 0{
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment, err := c.service.GetCommentById(ctx, uint64(commentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	userIdToken, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Cannot get user id from token"})
		return	
	}

	userId, ok := userIdToken.(uint32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: err.Error()})
		return	
	}

	if uint64(userId) != uint64(comment.UserId) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "does not have access to edit other user's data"})
		return
	}

	req := model.CommentRequest{}
	if err = ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment.Message = req.Message

	comment, err = c.service.UpdateComment(ctx, comment)
	if  err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "Success",
		Message: "OK",
		Data: comment,
	})
}

func (c commentHandlerImpl) DeleteComment(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil || commentId == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment, err := c.service.GetCommentById(ctx, uint64(commentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	userIdToken, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Please login first"})
		return
	}

	userId, ok := userIdToken.(uint32)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "cannot change the data type from interface to uint32"})
		return
	}

	if uint64(userId) != uint64(comment.UserId){
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "does not have access to edit other user's data"})
		return
	}

	if err = c.service.DeleteComment(ctx, uint64(commentId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "Success",
		Message: "Data has been deleted",
		Data: comment,
	})

}