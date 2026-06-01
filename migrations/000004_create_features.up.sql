CREATE TABLE features (
    id UUID PRIMARY KEY,

    environment_id UUID NOT NULL
        REFERENCES environments(id)
        ON DELETE CASCADE,

    key VARCHAR(255) NOT NULL,

    name VARCHAR(255) NOT NULL,

    description TEXT,

    enabled BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);