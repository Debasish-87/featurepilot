CREATE TABLE api_keys (
    id UUID PRIMARY KEY,

    organization_id UUID NOT NULL
        REFERENCES organizations(id)
        ON DELETE CASCADE,

    name VARCHAR(255) NOT NULL,

    key_hash VARCHAR(255) NOT NULL UNIQUE,

    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_api_keys_organization_id
ON api_keys(organization_id);