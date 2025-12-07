package api_auth_v1

import (
	apiauthv1workos "platform/backend/internal/api/auth/v1/workos"
	"platform/backend/pkg/workos"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	WorksOsClient workos.Client
}

func Register(authV1 *gin.RouterGroup, deps Dependencies) {
	handler := apiauthv1workos.NewWorkosHandler(deps.WorksOsClient)

	authV1.GET("/login", handler.Login)
	authV1.GET("/login/callback", handler.LoginCallback)
	authV1.GET("/logout", handler.Logout)
	authV1.GET("/logout/callback", handler.LogoutCallback)
}
