package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rifqidamarali/final-project-golang006/internal/model"
	"github.com/rifqidamarali/final-project-golang006/internal/service"
	"github.com/rifqidamarali/final-project-golang006/pkg"
)

type SocialMediaHandler interface {
	CreateSocialMedia(ctx *gin.Context)
	GetAllSocialMediaByUserId(ctx *gin.Context)
	GetSocialMediaById(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHandlerImpl struct{
	service service.SocialMediaServcie
}

func NewSocialMediaHandler(service service.SocialMediaServcie) SocialMediaHandler{
	return &socialMediaHandlerImpl{
		service: service,
	}
}

func(s *socialMediaHandlerImpl) CreateSocialMedia(ctx *gin.Context){
	req := model.SocialMediaRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.APIResponse{
			Status: "Failed",
			Message: "Failed",
			Errors: "Failed binding payload to struct",
		})
		return
	}
	// validate := validator.New()
	// if err := validate.Struct(req); err != nil {
	// 	var errorMessages []string
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed on the '%s' tag", err.Field(), err.Tag()))
	// }
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{
	// 		Message: "Failed",
	// 		Errors: errorMessages,
	// 	})
	// 	return
	// }
	errorStrings := model.ValidateSocialMediaRequest(req)
	if len(errorStrings) > 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Message: "check the payload",
			Errors: errorStrings,
		})
		return
	}

	userIdToken, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Please login first"})
		return
	}
	var userId uint64
	switch v := userIdToken.(type) {
		case uint64:
			userId = uint64(v)
		case uint32:
			userId = uint64(v) 
		case int:
			userId = uint64(v)
		case int64:
			userId = uint64(v) 
		default:
       		userId = 0	
	} 
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"user id not valid"},
		})
		return
	}
	res, err := s.service.CreateSocialMedia(ctx, req, userId)
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

func (s *socialMediaHandlerImpl) GetAllSocialMediaByUserId(ctx *gin.Context){
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if userId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Message: "failed",
			Errors: []string{"please input valid user id"},
		})
		return
	}

	socialMedias, err := s.service.GetAllSocialMediasByUserId(ctx, uint64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Message: "failed",
			Errors: []string{"internal server error"},
		})
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "Success",
		Message: "Success get user social medias",
		Data: socialMedias,
	})
}

func (s *socialMediaHandlerImpl) GetSocialMediaById(ctx *gin.Context){
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Message: "cannot find social media with that id",
			Errors: []string{err.Error()},
		})
		return
	}

	socialMedia, err := s.service.GetSocialMediaById(ctx, uint64(socialMediaId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
			Message: "Failed to get social media",
			Errors: ctx.Errors.Errors(),
		})
		return
	}	

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "success",
		Message: "Ok",
		Data: socialMedia,
	})
}

func (s *socialMediaHandlerImpl) UpdateSocialMedia(ctx *gin.Context){
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"social media id should be unsigned integer number"},
		})
		return
	}
	userIdToken, exists := ctx.Get("UserId")
	if !exists{
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"please login first"},
		})
	}
	var userId uint64
	switch v := userIdToken.(type) {
		case uint64:
			userId = uint64(v)
		case uint32:
			userId = uint64(v) 
		case int:
			userId = uint64(v)
		case int64:
			userId = uint64(v) 
		default:
       		userId = 0	
	} 
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"user id not valid"},
		})
		return
	}

	req := model.SocialMediaRequest{}
	if err = ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	input := model.SocialMediaUpdate{
		SocialMediaRequest: req,
		SocialMediaId: uint64(socialMediaId),
		UserId: userId,
		
	}

	socialMedia, err := s.service.UpdateSocialMedia(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			fmt.Println("ErrNotFound matched")
			ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrForbidden):
			fmt.Println("ErrForbidden matched")
			ctx.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: err.Error()})
		default:
			fmt.Printf("Unhandled error: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
				Message: "Failed",
				Errors:  []string{"internal server error"},
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, pkg.APIResponse{
		Status: "success",
		Message: fmt.Sprintf("success  update %v", socialMedia.Id),
		Data: socialMedia,
	})
}

func (s *socialMediaHandlerImpl) DeleteSocialMedia(ctx *gin.Context){
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil || socialMediaId == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"param should be uint64"},
		})
	}

	userIdToken, exists := ctx.Get("UserId")
	if !exists{
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"please login first"},
		})
	}
	var userId uint64
	switch v := userIdToken.(type) {
		case uint64:
			userId = uint64(v)
		case uint32:
			userId = uint64(v) 
		case int:
			userId = uint64(v)
		case int64:
			userId = uint64(v) 
		default:
       		userId = 0	
	} 
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "Failed",
			Errors: []string{"user id not valid"},
		})
		return
	}

	input := model.SocialMediaDelete{
		SocialMediaId: uint64(socialMediaId),
		UserId: uint64(userId),
	}

	err = s.service.DeleteSocialMedia(ctx, input)
	if err != nil {
		// fmt.Printf("Error returned: %v\n", err)
		switch {
		case errors.Is(err, service.ErrNotFound):
			fmt.Println("ErrNotFound matched")
			ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrForbidden):
			fmt.Println("ErrForbidden matched")
			ctx.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: err.Error()})
		default:
			fmt.Printf("Unhandled error: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
				Message: "Failed",
				Errors:  []string{"internal server error"},
			})
		}
		return
	}

ctx.JSON(http.StatusOK, pkg.APIResponse{
    Status:  "Success",
    Message: "Social media has been deleted",
})



}