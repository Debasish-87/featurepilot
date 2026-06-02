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
    'feature_' || (1 + floor(random()*1000)),
    md5(random()::text),
    random() > 0.2,
    NOW() - (random() * interval '365 days')
FROM generate_series(1,500000);