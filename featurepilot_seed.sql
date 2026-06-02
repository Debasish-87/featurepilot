-- FeaturePilot demo seed data

INSERT INTO organizations (id,name,created_at,updated_at)
VALUES (
gen_random_uuid(),
'FeaturePilot',
NOW(),
NOW()
)
ON CONFLICT DO NOTHING;

INSERT INTO projects (
id,
organization_id,
name,
created_at,
updated_at
)
SELECT
gen_random_uuid(),
id,
'Checkout Service',
NOW(),
NOW()
FROM organizations
WHERE name='FeaturePilot'
LIMIT 1;

INSERT INTO environments (
id,
project_id,
name,
created_at,
updated_at
)
SELECT
gen_random_uuid(),
id,
'production',
NOW(),
NOW()
FROM projects
WHERE name='Checkout Service'
LIMIT 1;

-- Features

INSERT INTO features (
id,
environment_id,
key,
name,
description,
enabled,
rollout_percentage,
created_at,
updated_at
)
SELECT
gen_random_uuid(),
e.id,
'checkout_redesign',
'Checkout Redesign',
'Modern checkout experience',
true,
100,
NOW(),
NOW()
FROM environments e
WHERE e.name='production';

INSERT INTO features (
id,
environment_id,
key,
name,
description,
enabled,
rollout_percentage,
created_at,
updated_at
)
SELECT
gen_random_uuid(),
e.id,
'new_dashboard',
'New Dashboard',
'Analytics dashboard rollout',
true,
75,
NOW(),
NOW()
FROM environments e
WHERE e.name='production';

INSERT INTO features (
id,
environment_id,
key,
name,
description,
enabled,
rollout_percentage,
created_at,
updated_at
)
SELECT
gen_random_uuid(),
e.id,
'ai_search',
'AI Search',
'Semantic search engine',
true,
50,
NOW(),
NOW()
FROM environments e
WHERE e.name='production';

-- Releases

INSERT INTO releases (
id,
project_id,
version,
status,
created_at
)
SELECT
gen_random_uuid(),
p.id,
'v1.0.0',
'ACTIVE',
NOW()
FROM projects p
LIMIT 1;

INSERT INTO releases (
id,
project_id,
version,
status,
created_at
)
SELECT
gen_random_uuid(),
p.id,
'v1.1.0',
'ACTIVE',
NOW()
FROM projects p
LIMIT 1;

-- Release Metrics

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
(random()*3)+1,
(random()*150)+80,
NOW()
FROM releases r;

-- Evaluation Events

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
WHEN random() < 0.33 THEN 'checkout_redesign'
WHEN random() < 0.66 THEN 'new_dashboard'
ELSE 'ai_search'
END,
md5(random()::text),
random() > 0.2,
NOW() - (random() * interval '30 days')
FROM generate_series(1,1000);
