package api_auth_v1_workos

import (
	authv1 "platform/backend/pkg/auth/v1"
	"platform/backend/pkg/workos"
)

const (
	SessionKeyAuthState = "workos_auth_state"
)

type workosHandler struct {
	client workos.Client
}

func NewWorkosHandler() authv1.AuthHandler {
	return &workosHandler{}
}
