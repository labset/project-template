package api_auth_v1_workos

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
)

func (w *workosHandler) Login(ctx *gin.Context) {
	state := uuid.Must(uuid.NewV7())

	session := sessions.Default(ctx)
	session.Set(SessionKeyAuthState, state.String())

	if err := session.Save(); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to save session"},
		)

		return
	}

	authURL, err := w.client.UserManagement().GetAuthorizationUrl(state.String())
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to get authorization URL"},
		)

		return
	}

	ctx.Redirect(http.StatusFound, authURL.String())
}
