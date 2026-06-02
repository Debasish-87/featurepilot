CREATE TABLE release_metrics (
    id UUID PRIMARY KEY,

    release_id UUID NOT NULL
        REFERENCES releases(id)
        ON DELETE CASCADE,

    error_rate NUMERIC NOT NULL,

    latency_ms NUMERIC NOT NULL,

    created_at TIMESTAMP NOT NULL
);