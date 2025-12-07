package workos

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/workos/workos-go/v5/pkg/organizations"
)

type Organisations interface {
	GetOrgByExternalID(
		ctx context.Context,
		externalID uuid.UUID,
	) (*organizations.Organization, error)
	CreateOrg(
		ctx context.Context,
		name string,
		externalID uuid.UUID,
	) (*organizations.Organization, error)
}

type organisations struct {
	wrapped *organizations.Client
}

func newOrganisations(cfg ClientConfig) Organisations {
	wrapped := organizations.DefaultClient
	wrapped.APIKey = cfg.APIKey

	return &organisations{
		wrapped: wrapped,
	}
}

func (o *organisations) GetOrgByExternalID(
	ctx context.Context,
	externalID uuid.UUID,
) (*organizations.Organization, error) {
	response, err := o.wrapped.GetOrganizationByExternalID(
		ctx,
		organizations.GetOrganizationByExternalIDOpts{
			ExternalID: externalID.String(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization by external ID from WorkOS: %w", err)
	}

	return &response, nil
}

func (o *organisations) CreateOrg(
	ctx context.Context,
	name string,
	externalID uuid.UUID,
) (*organizations.Organization, error) {
	response, err := o.wrapped.CreateOrganization(ctx, organizations.CreateOrganizationOpts{
		Name:       name,
		ExternalID: externalID.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create organization in WorkOS: %w", err)
	}

	return &response, nil
}
