-- +goose Up
CREATE TYPE tenant_membership_role AS ENUM ('admin', 'member');
CREATE TYPE tenant_membership_status AS ENUM ('provisioning', 'active', 'suspended');

CREATE TABLE tenant_membership
(
    id          UUID PRIMARY KEY,
    identity_id UUID NOT NULL REFERENCES platform_identity (id) ON DELETE CASCADE,
    tenant_id   UUID NOT NULL REFERENCES tenant (id) ON DELETE CASCADE,
    role        tenant_membership_role   NOT NULL DEFAULT 'member',
    status      tenant_membership_status NOT NULL DEFAULT 'provisioning',
    created_at  TIMESTAMP WITH TIME ZONE          DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT NOW(),
    deleted_at  TIMESTAMP WITH TIME ZONE,
    UNIQUE (identity_id, tenant_id)
);

-- Performance indexes
CREATE INDEX idx_tenant_membership_identity ON tenant_membership (identity_id);
CREATE INDEX idx_tenant_membership_tenant ON tenant_membership (tenant_id);
CREATE INDEX idx_tenant_membership_role ON tenant_membership (role);
CREATE INDEX idx_tenant_membership_status ON tenant_membership (status);

-- +goose Down
DROP INDEX IF EXISTS idx_tenant_membership_identity;
DROP INDEX IF EXISTS idx_tenant_membership_tenant;
DROP INDEX IF EXISTS idx_tenant_membership_role;
DROP INDEX IF EXISTS idx_tenant_membership_status;
DROP TABLE IF EXISTS tenant_membership;
DROP TYPE IF EXISTS tenant_membership_role;
DROP TYPE IF EXISTS tenant_membership_status;