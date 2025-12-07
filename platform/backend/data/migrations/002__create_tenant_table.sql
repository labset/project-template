-- +goose Up
CREATE TYPE tenant_type AS ENUM ('personal', 'team', 'enterprise');
CREATE TYPE tenant_status AS ENUM ('provisioning', 'active', 'suspended');

CREATE TABLE tenant
(
    id               UUID PRIMARY KEY,

    provider         platform_identity_provider NOT NULL,
    external_id      VARCHAR UNIQUE             NOT NULL,
    external_profile JSONB                      NOT NULL DEFAULT '{}',

    type             tenant_type                NOT NULL DEFAULT 'personal',
    status           tenant_status              NOT NULL DEFAULT 'provisioning',
    created_at       TIMESTAMP WITH TIME ZONE            DEFAULT NOW(),
    updated_at       TIMESTAMP WITH TIME ZONE            DEFAULT NOW(),
    deleted_at       TIMESTAMP WITH TIME ZONE
);

-- Performance indexes
CREATE INDEX idx_tenant_type ON tenant (type);
CREATE INDEX idx_tenant_status ON tenant (status);

-- +goose Down
DROP INDEX IF EXISTS idx_tenant_type;
DROP INDEX IF EXISTS idx_tenant_status;
DROP TABLE IF EXISTS tenant;
DROP TYPE IF EXISTS tenant_type;
DROP TYPE IF EXISTS tenant_status;