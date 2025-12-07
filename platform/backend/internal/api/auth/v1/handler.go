package api_auth_v1

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Login(ctx *gin.Context)
	LoginCallback(ctx *gin.Context)
	Logout(ctx *gin.Context)
	LogoutCallback(ctx *gin.Context)
}

type authHandler struct{}

func newAuthHandler() AuthHandler {
	return &authHandler{}
}
