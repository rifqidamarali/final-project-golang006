package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/final-project-golang006/internal/handler"
	"github.com/rifqidamarali/final-project-golang006/internal/middleware"
)

type CommentRouter interface {
	Mount()
}

type CommentRouterImpl struct {
	v *gin.RouterGroup
	handler handler.CommentHandler
	auth middleware.Authorization
}

func NewCommentRouter(v *gin.RouterGroup, handler handler.CommentHandler, auth middleware.Authorization) CommentRouter{
	return &CommentRouterImpl{
		v : v,
		handler: handler,
		auth: auth,
	}
}

func (c *CommentRouterImpl) Mount(){
	c.v.GET("/photos/:photoId/comments", c.handler.GetAllCommentsByPhotoId)
	c.v.GET("/comments/:commentId", c.handler.GetCommentById)
	

	c.v.Use(c.auth.CheckAuth)
	c.v.POST("/photos/:photoId/comments", c.handler.CreateComment)
	c.v.PUT("/comments/:commentId", c.handler.UpdateComment)
	c.v.DELETE("/comments/:commentId", c.handler.DeleteComment)

}