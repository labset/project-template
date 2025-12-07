package api_auth_v1_workos

import (
	authv1 "platform/backend/pkg/auth/v1"
)

type workosHandler struct{}

func NewWorkosHandler() authv1.AuthHandler {
	return &workosHandler{}
}
