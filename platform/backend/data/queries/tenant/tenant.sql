-- name: GetTenantByID :one
SELECT id, provider, external_id, external_profile, type, status, created_at, updated_at
FROM tenant
WHERE id = @id;


-- name: CreateTenant :exec
INSERT INTO tenant (
    id,
    provider,
    external_id,
    type,
    status,
    created_at,
    updated_at
)
VALUES (
           @id,
           @provider,
           @external_id,
          @type,
           @status,
           NOW(),
           NOW()
       );