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
	v       *gin.RouterGroup
	// handler handler.SocialMedia
	auth    middleware.Authorization
}

func NewSocialMediaRouter(v *gin.RouterGroup, handler handler.SocialMediaHandler, auth middleware.Authorization) SocialMediaRouter {
	return &socialMediaRouterImpl{v: v, handler: handler, auth: auth}
}

func (s *socialMediaRouterImpl) Mount() {
	s.v.Use(s.auth.CheckAuth)
	s.v.POST("", s.handler.PostSocialMedia)
	s.v.GET("", s.handler.GetAllSocialMediasByUserId)
	s.v.GET("/:id", s.handler.GetSocialMediaById)
	s.v.PUT("/:id", s.handler.UpdateSocialMedia)
	s.v.DELETE("/:id", s.handler.DeleteSocialMedia)
}