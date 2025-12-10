package workos

import (
	"context"
	"fmt"
	"net/url"

	"github.com/workos/workos-go/v5/pkg/usermanagement"
)

type UserManagement interface {
	GetLogoutURL(sessionID string) (*url.URL, error)
	GetAuthorizationURL(state string) (*url.URL, error)
	AuthenticateWithCode(
		ctx context.Context,
		code string,
	) (*usermanagement.AuthenticateResponse, error)
	GetOrgMembership(
		ctx context.Context,
		opts usermanagement.GetOrganizationMembershipOpts,
	) (*usermanagement.OrganizationMembership, error)
	CreateOrgMembership(
		ctx context.Context,
		opts usermanagement.CreateOrganizationMembershipOpts,
	) (*usermanagement.OrganizationMembership, error)
	GetJWKSURL() (*url.URL, error)
}

type userManagement struct {
	clientID    string
	apiKey      string
	redirectURI string

	wrapped *usermanagement.Client
}

func newUserManagement(cfg ClientConfig) UserManagement {
	return &userManagement{
		clientID:    cfg.ClientID,
		apiKey:      cfg.APIKey,
		redirectURI: cfg.RedirectURI,
		wrapped:     usermanagement.NewClient(cfg.APIKey),
	}
}

func (u *userManagement) GetJWKSURL() (*url.URL, error) {
	return u.wrapped.GetJWKSURL(u.clientID)
}

func (u *userManagement) GetAuthorizationURL(state string) (*url.URL, error) {
	return u.wrapped.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID:    u.clientID,
		Provider:    "authkit",
		RedirectURI: u.redirectURI,
		State:       state,
	})
}

func (u *userManagement) GetLogoutURL(sessionID string) (*url.URL, error) {
	return u.wrapped.GetLogoutURL(usermanagement.GetLogoutURLOpts{
		SessionID: sessionID,
	})
}

func (u *userManagement) AuthenticateWithCode(
	ctx context.Context,
	code string,
) (*usermanagement.AuthenticateResponse, error) {
	response, err := u.wrapped.AuthenticateWithCode(ctx, usermanagement.AuthenticateWithCodeOpts{
		Code:     code,
		ClientID: u.clientID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user with WorkOS: %w", err)
	}

	return &response, nil
}

func (u *userManagement) GetOrgMembership(
	ctx context.Context,
	opts usermanagement.GetOrganizationMembershipOpts,
) (*usermanagement.OrganizationMembership, error) {
	response, err := u.wrapped.GetOrganizationMembership(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get org membership from WorkOS: %w", err)
	}

	return &response, nil
}

func (u *userManagement) CreateOrgMembership(
	ctx context.Context,
	opts usermanagement.CreateOrganizationMembershipOpts,
) (*usermanagement.OrganizationMembership, error) {
	response, err := u.wrapped.CreateOrganizationMembership(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create org membership in WorkOS: %w", err)
	}

	return &response, nil
}
