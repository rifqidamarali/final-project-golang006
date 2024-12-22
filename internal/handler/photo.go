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

type PhotoHandler interface {
	CreatePhoto(ctx *gin.Context)
	GetAllPhotosById(ctx *gin.Context)
	GetPhotoById(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
}

type photoHandlerImpl struct {
	service service.PhotoService
}

func NewPhotoHandler(service service.PhotoService) PhotoHandler {
	return &photoHandlerImpl {
		service: service,
	}
}

func (u *photoHandlerImpl) CreatePhoto(ctx *gin.Context){
	photoRequest := model.PhotoRequest{}
	if err := ctx.Bind(&photoRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid request parameter"})
	}	

	validate := validator.New()
	err := validate.Struct(photoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid validating photo"})
		return
	}
	userId, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: err.Error()})
		return
	}
    // Add UserId to a standard context.Context
    stdCtx := context.WithValue(ctx.Request.Context(), "UserId", userId)
	res, err := u.service.CreatePhoto(stdCtx, photoRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (u *photoHandlerImpl) GetAllPhotosById(ctx *gin.Context) {
	// get user id
	id, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param get users photo"})
		return
	}

	photos, err := u.service.GetAllPhotosById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{Status: "Success", Message: "OK", Data: photos})
}

func (u *photoHandlerImpl) GetPhotoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photo, err := u.service.GetPhotoById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if photo.Id == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "photo doesn't exist"})
		return
	}
	ctx.JSON(http.StatusOK, pkg.APIResponse{Status: "Success", Message: "OK", Data: photo})

}

func (u *photoHandlerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	} 

	userIdPhoto, err := u.service.GetPhotoById(ctx, uint64(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
	}

	userIdToken, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Cannot get user id from token"})
		return	
	}

	uint32Value, ok := userIdToken.(uint32)
    if !ok {
        ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "cannot change the data type from interface to uint32"})
		return
    }

	if uint64(userIdPhoto.UserId) != uint64(uint32Value) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "does not have access to edit other user's data"})
		return
	}

	photo, err := u.service.DeletePhoto(ctx, uint64(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	} 

	if photo.Id == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "photo doesn't exist"})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{Status: "Success", Message: "photo has been deleted"})
}

func (u *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil  || photoId == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
	}
	Photo, err := u.service.GetPhotoById(ctx, uint64(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	userIdToken, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Cannot get user id from token"})
		return	
	}

	uint32Value, ok := userIdToken.(uint32)
    if !ok {
        ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "cannot change the data type from interface to uint32"})
		return
    }

	if uint64(Photo.UserId) != uint64(uint32Value) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "does not have access to edit other user's data"})
		return
	}
	
	req := model.PhotoRequest{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photoData := model.Photo{
		Id: uint64(photoId),
		Caption: req.Caption,
		Title: req.Title,
		Url: req.Url,
		
	}

	photo, err := u.service.UpdatePhoto(ctx, photoData)
	if  err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "Success",
		Message: "OK",
		Data: photo,
	})
}