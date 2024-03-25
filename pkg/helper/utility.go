package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromGinCtx(ctx *gin.Context) (uint32, error) {
	userIdRaw, isExist := ctx.Get("UserId")
	if !isExist {
		return 0, errors.New("cannot get payload in access token")
	}

	userIdFloat := userIdRaw.(float64)
	userId := int(userIdFloat)
	if userId == 0 {
		return 0, errors.New("cannot get payload in access token")
	}

	return uint32(userId), nil
}