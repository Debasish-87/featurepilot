CREATE TABLE environments (
    id UUID PRIMARY KEY,

    project_id UUID NOT NULL
        REFERENCES projects(id),

    name VARCHAR(255) NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
