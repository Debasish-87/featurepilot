CREATE TABLE releases (
    id UUID PRIMARY KEY,

    project_id UUID NOT NULL
        REFERENCES projects(id)
        ON DELETE CASCADE,

    version VARCHAR(100) NOT NULL,

    status VARCHAR(50) NOT NULL,

    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_releases_project
ON releases(project_id);
