-- +goose Up
CREATE TYPE platform_identity_status AS ENUM ('provisioning', 'active', 'suspended');
CREATE TYPE platform_identity_provider AS ENUM ('workos', 'auth0');

CREATE TABLE platform_identity
(
    id               UUID PRIMARY KEY,

    provider         platform_identity_provider NOT NULL,
    external_id      VARCHAR UNIQUE             NOT NULL,
    external_profile JSONB                      NOT NULL DEFAULT '{}',

    status           platform_identity_status   NOT NULL DEFAULT 'provisioning',
    created_at       TIMESTAMP WITH TIME ZONE            DEFAULT NOW(),
    updated_at       TIMESTAMP WITH TIME ZONE            DEFAULT NOW(),
    deleted_at       TIMESTAMP WITH TIME ZONE
);

-- Performance indexes
CREATE INDEX idx_platform_identity_provider ON platform_identity (provider, external_id);

-- +goose Down
DROP INDEX IF EXISTS idx_platform_identity_provider;
DROP TABLE IF EXISTS platform_identity;
DROP TYPE IF EXISTS platform_identity_status;