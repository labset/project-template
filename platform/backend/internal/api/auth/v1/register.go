package api_auth_v1

import "github.com/gin-gonic/gin"

type Dependencies struct {
}

func Register(authV1 *gin.RouterGroup, deps Dependencies) {
	handler := newAuthHandler()

	authV1.GET("/login", handler.Login)
	authV1.GET("/login/callback", handler.LoginCallback)
	authV1.GET("/logout", handler.Logout)
	authV1.GET("/logout/callback", handler.LogoutCallback)
}
