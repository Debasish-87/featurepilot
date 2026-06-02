-- 50,000 evaluation events

INSERT INTO evaluation_events (
    id,
    environment,
    feature_key,
    user_id,
    enabled,
    created_at
)
SELECT
    gen_random_uuid(),
    'production',
    CASE
        WHEN random() < 0.20 THEN 'checkout_redesign'
        WHEN random() < 0.40 THEN 'new_dashboard'
        WHEN random() < 0.60 THEN 'ai_search'
        WHEN random() < 0.80 THEN 'dark_mode'
        ELSE 'payment_v2'
    END,
    md5(random()::text),
    random() > 0.2,
    NOW() - (random() * interval '90 days')
FROM generate_series(1,50000);

-- 500 release metrics

INSERT INTO release_metrics (
    id,
    release_id,
    error_rate,
    latency_ms,
    created_at
)
SELECT
    gen_random_uuid(),
    r.id,
    random() * 20,
    50 + random() * 1000,
    NOW() - (random() * interval '30 days')
FROM releases r
CROSS JOIN generate_series(1,100);

-- 10 extra releases

INSERT INTO releases (
    id,
    project_id,
    version,
    status,
    created_at
)
SELECT
    gen_random_uuid(),
    '2203f63a-9e27-4dd6-af96-3ef4f56721bb',
    'v2.' || gs || '.0',
    CASE
        WHEN random() < 0.5 THEN 'ACTIVE'
        WHEN random() < 0.8 THEN 'ROLLING_OUT'
        ELSE 'FAILED'
    END,
    NOW() - (random() * interval '60 days')
FROM generate_series(1,10) gs;