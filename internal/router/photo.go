package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/final-project-golang006/internal/handler"
	"github.com/rifqidamarali/final-project-golang006/internal/middleware"
)

type PhotoRouter interface {
	Mount()
}

type photoRouterImpl struct {
	v 	*gin.RouterGroup
	handler handler.PhotoHandler
	auth middleware.Authorization
}

func NewPhotoRouter (v *gin.RouterGroup, handler handler.PhotoHandler, auth middleware.Authorization) PhotoRouter{
	return &photoRouterImpl{
		v : v, 
		handler: handler, 
		auth: auth}
}

func (u *photoRouterImpl) Mount() {
	u.v.GET("/users/:userId/photos", u.handler.GetAllPhotosById)
	u.v.GET("/photos/:photoId", u.handler.GetPhotoById)
	
	u.v.Use(u.auth.CheckAuth)
	u.v.POST("/photos", u.handler.CreatePhoto)
	u.v.DELETE("/photos/:photoId", u.handler.DeletePhoto)
	u.v.PUT("/photos/:photoId", u.handler.UpdatePhoto)
}