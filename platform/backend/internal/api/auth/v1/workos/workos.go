package api_auth_v1_workos

import (
	"github.com/gin-gonic/gin"
	authv1 "platform/backend/pkg/auth/v1"
)

const (
	SessionKeyAuthState = "workos_auth_state"
)

type workosHandler struct {
}

func NewWorkosHandler() authv1.AuthHandler {
	return &workosHandler{}
}
