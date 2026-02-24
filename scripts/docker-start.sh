#!/bin/bash
# SWJTU CTF OJ - Docker Infrastructure Start Script

set -e

echo "🚀 Starting SWJTU CTF OJ infrastructure..."

# Check if containers are already running
if docker ps --format '{{.Names}}' | grep -q ctfoj-db; then
    echo "✅ PostgreSQL already running"
else
    echo "📦 Starting PostgreSQL..."
    docker run -d --name ctfoj-db \
      -e POSTGRES_DB=ctfoj \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=postgres \
      -p 5432:5432 \
      postgres:15-alpine
fi

if docker ps --format '{{.Names}}' | grep -q ctfoj-redis; then
    echo "✅ Redis already running"
else
    echo "📦 Starting Redis..."
    docker run -d --name ctfoj-redis \
      -p 6379:6379 \
      redis:7-alpine
fi

# Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready..."
sleep 3

# Test connection
if docker exec ctfoj-db psql -U postgres -d ctfoj -c "SELECT 1;" > /dev/null 2>&1; then
    echo "✅ PostgreSQL is ready"
else
    echo "❌ PostgreSQL connection failed"
    exit 1
fi

echo ""
echo "🎉 Infrastructure started successfully!"
echo ""
echo "Services:"
echo "  - PostgreSQL: localhost:5432 (db: ctfoj, user: postgres, pass: postgres)"
echo "  - Redis: localhost:6379"
echo ""
echo "To stop: docker stop ctfoj-db ctfoj-redis"
echo "To remove: docker rm ctfoj-db ctfoj-redis"
