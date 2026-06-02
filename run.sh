#!/bin/bash

echo "Starting FeaturePilot..."

docker compose up -d

echo ""
echo "FeaturePilot Ready"
echo ""
echo "Dashboard : http://localhost:3000"
echo "Analytics : http://localhost:8000/docs"
echo "API       : http://localhost:8080/health"
echo ""
echo "Status:"
docker ps