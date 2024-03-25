package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/final-project-golang006/internal/handler"
	"github.com/rifqidamarali/final-project-golang006/internal/middleware"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
	auth middleware.Authorization
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler, auth middleware.Authorization) UserRouter {
	return &userRouterImpl{v: v, handler: handler, auth: auth}
}

func (u *userRouterImpl) Mount() {
	// activity
	// /users/sign-up
	u.v.POST("/sign-up", u.handler.UserSignUp)
	u.v.POST("/sign-in", u.handler.UserSignIn)
	u.v.GET("/:id", u.handler.GetUserById)
	// users
	u.v.Use(u.auth.CheckAuth)
	// /users
	// u.v.GET("", u.handler.GetUsers)
	// /users/:id
	u.v.PUT("/:id", u.handler.UserUpdate)
	u.v.DELETE("/:id", u.handler.DeleteUserById)
	
}
