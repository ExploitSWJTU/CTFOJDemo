#!/bin/bash
# SWJTU CTF OJ - Docker Infrastructure Stop Script

echo "🛑 Stopping SWJTU CTF OJ infrastructure..."

docker stop ctfoj-db ctfoj-redis 2>/dev/null && echo "✅ Containers stopped"
echo "To start again: ./scripts/docker-start.sh"
