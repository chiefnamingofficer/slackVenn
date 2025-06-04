#!/bin/bash

set -e

# Load environment variables
if [ -f "env/.env" ]; then
    set -a
    source env/.env
    set +a
    echo "✅ Environment loaded from env/.env"
else
    echo "⚠️  No env/.env file found. Using system environment variables."
    echo "💡 Copy env/.env.example to env/.env and add your SLACK_TOKEN"
fi

# --- Handle CLI flags ---
DRY_RUN=false
POSITIONAL=()

while [[ $# -gt 0 ]]; do
  case "$1" in
    --dry-run)
      DRY_RUN=true
      shift
      ;;
    -*)
      echo "❌ Unknown flag: $1"
      exit 1
      ;;
    *)
      POSITIONAL+=("$1")
      shift
      ;;
  esac
done

set -- "${POSITIONAL[@]}"

# --- Collect channel IDs ---
if [ "$#" -eq 2 ]; then
  CHANNEL_A="$1"
  CHANNEL_B="$2"
else
  echo "ℹ️  Channel IDs not provided. Enter them interactively."
  read -rp "🔹 Channel A ID: " CHANNEL_A
  read -rp "🔹 Channel B ID: " CHANNEL_B
fi

# --- Output setup ---
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
OUTDIR="history"
FILENAME="channel_comparison_${TIMESTAMP}.csv"
FULLPATH="$OUTDIR/$FILENAME"
SYMLINK="CURRENT"

echo "📊 slackVenn: Comparing $CHANNEL_A vs $CHANNEL_B"
echo "📁 Output file: $FULLPATH"
echo "🔗 Symlink: $SYMLINK → $FULLPATH"

if $DRY_RUN; then
  echo "✅ Dry run complete. No actions taken."
  exit 0
fi

# --- Create output dir and run the comparison ---
mkdir -p "$OUTDIR"

(
  echo "Status,Username"
  go run main.go "$CHANNEL_A" "$CHANNEL_B" | \
  awk '
    /Users in BOTH/ {section="In Both"; next}
    /Users ONLY in Channel A/ {section="Only in Channel A"; next}
    /Users ONLY in Channel B/ {section="Only in Channel B"; next}
    /^ -/ {gsub(/^ - /, "", $0); print section "," $0}
  ' | sort -t, -k1,1 -k2,2
) > "$FULLPATH"

ln -sf "$FULLPATH" "$SYMLINK"

echo "✅ CSV written to $FULLPATH"
echo "🔗 CURRENT → $FULLPATH" 