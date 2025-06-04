# slackVenn Setup Guide

Complete step-by-step instructions for setting up slackVenn to analyze Slack channel membership overlaps.

## 🎯 Quick Start

```bash
# 1. Clone and setup
git clone git@github.com:chiefnamingofficer/slackVenn.git
cd slackVenn

# 2. Install dependencies
go mod tidy

# 3. Configure your Slack token
cp env/.env.example env/.env
# Edit env/.env with your Slack token

# 4. Test the setup
./slackVenn.sh --dry-run C1234567890 C0987654321
```

## 📋 Prerequisites

- **Go 1.21+** - [Download Go](https://golang.org/dl/)
- **Slack Workspace Access** - Admin or app management permissions
- **Basic terminal/command line knowledge**

## 🔑 Getting a Slack Bot Token

### Step 1: Create a Slack App

1. Go to [Slack App Dashboard](https://api.slack.com/apps)
2. Click **"Create New App"**
3. Choose **"From scratch"**
4. Enter App Name: `slackVenn Channel Analyzer`
5. Select your workspace
6. Click **"Create App"**

### Step 2: Configure OAuth Scopes

1. In your app dashboard, go to **"OAuth & Permissions"**
2. Scroll down to **"Scopes"**
3. Under **"Bot Token Scopes"**, add these permissions:

**Required Scopes:**
- `channels:read` - View basic information about public channels
- `groups:read` - View basic information about private channels  
- `users:read` - View people in the workspace
- `users:read.email` - View email addresses of people (optional)

### Step 3: Install App to Workspace

1. Scroll up to **"OAuth Tokens for Your Workspace"**
2. Click **"Install to Workspace"**
3. Review permissions and click **"Allow"**
4. Copy the **"Bot User OAuth Token"** (starts with `xoxb-`)

### Step 4: Configure slackVenn

1. **Copy the environment template:**
   ```bash
   cp env/.env.example env/.env
   ```

2. **Edit the .env file:**
   ```bash
   # Open with your preferred editor
   nano env/.env
   # or
   vim env/.env
   # or
   code env/.env
   ```

3. **Add your token:**
   ```bash
   SLACK_TOKEN=xoxb-your-slack-bot-token-here
   ```

## 🚀 Usage

### Loading Environment

Always load your environment before using slackVenn:

```bash
# Load environment variables
source scripts/load-env.sh
```

### Finding Channel IDs

**Method 1: From Slack URL**
1. Open the channel in Slack
2. Copy URL: `https://yourworkspace.slack.com/archives/C1234567890`
3. Channel ID is: `C1234567890`

**Method 2: Using Slack API**
```bash
# List all channels (requires environment loaded)
curl -H "Authorization: Bearer $SLACK_TOKEN" \
  "https://slack.com/api/conversations.list?types=public_channel,private_channel" | jq '.channels[] | {name: .name, id: .id}'
```

### Running Comparisons

**Using the shell script (recommended):**
```bash
# Interactive mode
./slackVenn.sh

# With channel IDs
./slackVenn.sh C1234567890 C0987654321

# Dry run (test without creating files)
./slackVenn.sh --dry-run C1234567890 C0987654321
```

**Using Go directly:**
```bash
go run main.go C1234567890 C0987654321
```

**Using compiled binary:**
```bash
go build -o slackVenn main.go
./slackVenn C1234567890 C0987654321
```

## 📁 Project Structure

```
slackVenn/
├── env/                      # Environment configuration
│   ├── .env.example         #   Template for environment variables
│   └── .env                 #   Your actual environment (gitignored)
├── scripts/                 # Utility scripts
│   ├── load-env.sh         #   Environment loader script
│   └── run-tests.sh        #   Test suite runner
├── docs/                    # Documentation
│   └── SETUP.md            #   This setup guide
├── tests/                   # Test suite
│   ├── main_test.go        #   Unit tests
│   ├── mock_test.go        #   Mock tests
│   ├── README.md           #   Test documentation
│   └── results/            #   Test outputs (gitignored)
│       ├── .gitkeep        #     Preserves directory structure
│       ├── *_YYYYMMDD_HHMMSS.*  #  Timestamped test files
│       └── latest-*        #     Symlinks to latest results
├── history/                 # Generated comparison results (gitignored)
│   ├── .gitkeep            #   Preserves directory structure
│   └── *.csv               #   Timestamped comparison files
├── main.go                  # Core comparison logic
├── slackVenn.sh            # Shell script wrapper
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── .gitignore              # Git ignore rules
├── CURRENT                 # Symlink to latest result (gitignored)
└── README.md               # Main project documentation
```

## 🔒 Security Notes

- **Never commit your `.env` file** - It's already in `.gitignore`
- **Keep your Slack token private** - Don't share it in chat, logs, or screenshots
- **Use workspace-specific tokens** - Create separate apps for different workspaces
- **Rotate tokens regularly** - Generate new tokens if compromised

## 🐛 Troubleshooting

### "SLACK_TOKEN env var is required"
```bash
# Make sure you've loaded the environment
source scripts/load-env.sh

# Check if token is set
echo $SLACK_TOKEN
```

### "invalid_auth" error
- Verify your token is correct in `env/.env`
- Check that your app is installed in the workspace
- Ensure the token starts with `xoxb-`

### "Error getting members of channel"
- Verify channel IDs are correct (start with `C`)
- Ensure your bot has access to both channels
- Check that required scopes are added to your app

### **🔐 Private Channel Access Issues**

**For private channels, the slackVenn app must be invited to each channel:**

1. **Invite the app to private channels:**
   ```
   # In each private channel, type:
   /invite @slackVenn Channel Analyzer
   ```

2. **Alternative method via channel settings:**
   - Go to the private channel
   - Click the channel name → Settings → Integrations
   - Click "Add apps" → Search for "slackVenn Channel Analyzer"
   - Click "Add"

3. **Verify app access:**
   - The app should appear in the channel member list
   - You should see a message like "slackVenn Channel Analyzer was added to this channel"

**Common private channel errors:**
- `not_in_channel` - App not invited to private channel
- `channel_not_found` - App doesn't have access to view the channel
- `access_denied` - Insufficient permissions for private channel

**📝 Note:** Public channels don't require explicit invitation - the app can access them with proper scopes.

### Permission errors
```bash
# Make scripts executable
chmod +x slackVenn.sh
chmod +x scripts/load-env.sh
chmod +x scripts/run-tests.sh
```

### Test failures
```bash
# Run the test suite to verify everything works
./scripts/run-tests.sh
```

## 🧪 Testing

slackVenn includes a comprehensive test suite with timestamped results:

```bash
# Run all tests with coverage and benchmarks
./scripts/run-tests.sh

# Run specific test types
go test ./tests/                    # Unit tests only
go test -bench=. ./tests/          # Benchmarks only
go test -v -run TestMock ./tests/  # Mock tests only
```

**📊 Test features:**
- Unit tests with edge case coverage
- Mock tests with realistic Slack data
- Performance benchmarks for large channels
- Integration tests with dry-run validation
- Timestamped results prevent overwriting
- HTML coverage reports with line-by-line analysis

**📖 Complete testing guide: [tests/README.md](../tests/README.md)**

## 🤝 Sharing with Team

When sharing this project:

1. **Don't include your `.env` file**
2. **Share the repository without tokens**
3. **Point team members to this setup guide**
4. **Each person needs their own Slack app/token**

## 💡 Pro Tips

- **Bookmark useful channel IDs** in your `.env` file as comments
- **Use descriptive output filenames** by running comparisons for specific purposes
- **Set up aliases** for frequently compared channel pairs
- **Create team-specific documentation** with your common channel IDs
- **Run tests regularly** to ensure everything works: `./scripts/run-tests.sh`

---

**Questions?** Check the main [README.md](../README.md) for more details or open an issue. 