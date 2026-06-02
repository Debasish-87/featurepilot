CREATE TABLE evaluation_events (
    id UUID PRIMARY KEY,

    environment VARCHAR(255) NOT NULL,

    feature_key VARCHAR(255) NOT NULL,

    user_id VARCHAR(255) NOT NULL,

    enabled BOOLEAN NOT NULL,

    created_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_eval_feature
ON evaluation_events(feature_key);

CREATE INDEX idx_eval_created
ON evaluation_events(created_at DESC);