#!/bin/bash

# Load environment variables for slackVenn
# Usage: source scripts/load-env.sh

ENV_FILE="env/.env"

if [ ! -f "$ENV_FILE" ]; then
    echo "❌ Environment file not found: $ENV_FILE"
    echo "💡 Copy env/.env.example to env/.env and fill in your Slack token"
    return 1
fi

# Load environment variables
set -a  # automatically export all variables
source "$ENV_FILE"
set +a

echo "✅ Environment loaded from $ENV_FILE"
echo "📊 SLACK_TOKEN: ${SLACK_TOKEN:0:12}..." # Show first 12 chars only 