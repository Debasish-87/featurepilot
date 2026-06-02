#!/bin/sh

echo "Running migrations..."

migrate \
-path /app/migrations \
-database "postgres://postgres:postgres@postgres:5432/featurepilot?sslmode=disable" \
up || true

echo "Starting FeaturePilot..."

exec ./featurepilot