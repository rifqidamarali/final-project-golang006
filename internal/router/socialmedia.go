package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/final-project-golang006/internal/handler"
	"github.com/rifqidamarali/final-project-golang006/internal/middleware"
)

type SocialMediaRouter interface {
	Mount()
}

type socialMediaRouterImpl struct {
	v *gin.RouterGroup
	handler handler.SocialMediaHandler
	auth middleware.Authorization
}

func NewSocialMediaRouter(v *gin.RouterGroup, handler handler.SocialMediaHandler, auth middleware.Authorization) SocialMediaRouter{
	return &socialMediaRouterImpl{
		v : v,
		handler: handler,
		auth: auth,
	}
}

func (s *socialMediaRouterImpl) Mount(){
	s.v.GET("socialmedias/:socialMediaId", s.handler.GetSocialMediaById)
	s.v.GET("users/:userId/socialmedias", s.handler.GetAllSocialMediaByUserId)

	s.v.Use(s.auth.CheckAuth)
	s.v.POST("socialmedias", s.handler.CreateSocialMedia)
	s.v.DELETE("socialmedias/:socialMediaId", s.handler.DeleteSocialMedia)
	s.v.PUT("socialmedias/:socialMediaId", s.handler.UpdateSocialMedia)
	


}