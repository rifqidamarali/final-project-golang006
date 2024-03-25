package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/final-project-golang006/internal/service"
	"github.com/rifqidamarali/final-project-golang006/pkg"
	"github.com/rifqidamarali/final-project-golang006/pkg/helper"
)

type Authorization interface {
	CheckAuth(ctx *gin.Context)
}

type authorizationImpl struct {
	userService service.UserService
}

func NewAuthorization(userService service.UserService) Authorization {
	return &authorizationImpl{userService: userService}
}
func (a *authorizationImpl) CheckAuth(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	authArr := strings.Split(auth, " ")
	if len(authArr) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "unauthorized",
			Errors:  []string{"invalid token"},
		})
		return
	}
	if authArr[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "unauthorized",
			Errors:  []string{"invalid authorization method"},
		})
		return
	}

	token := authArr[1]
	claims, err := helper.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "unauthorized",
			Errors:  []string{"invalid token", "failed to decode"},
		})
		return
	}

	ctx.Set("UserId", claims["user_id"])

	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Failed when get user id"})
		return
	}

	user, err := a.userService.GetUserById(ctx, uint64(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "failed when get user id"})
		return
	}

	if user.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "user does not exist"})
		return
	}

	ctx.Next()
}
