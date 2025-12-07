package api_auth_v1_auth0

import (
	authv1 "platform/backend/pkg/auth/v1"
)

type auth0Handler struct{}

func NewAuth0Handler() authv1.AuthHandler {
	return &auth0Handler{}
}
