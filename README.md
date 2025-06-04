# slackVenn

**A Slack channel membership analyzer that creates Venn diagrams of user overlaps**

Visualize the relationships between Slack channels by finding who's in both, who's unique to each channel, and generating timestamped CSV reports for tracking changes over time.

## 🚀 Quick Start

```bash
# 1. Clone and setup
git clone <repository-url>
cd slackVenn

# 2. Install dependencies
go mod tidy

# 3. Configure your Slack token
cp env/.env.example env/.env
# Edit env/.env with your Slack token

# 4. Test the setup
./slackVenn.sh --dry-run C1234567890 C0987654321
```

## 🚀 Features

- ✅ **Channel Overlap Analysis** - Find users in both channels
- ✅ **Unique Member Detection** - Identify users exclusive to each channel  
- ✅ **Large Channel Support** - Handles 1000+ members with pagination
- ✅ **CSV Export** - Timestamped reports with historical tracking
- ✅ **Interactive Mode** - Prompts for channel IDs when not provided
- ✅ **Dry Run Support** - Test without generating files
- ✅ **Environment Management** - Secure token storage with .env files
- ✅ **Visual Output** - Color-coded emoji indicators for easy scanning
- ✅ **Comprehensive Testing** - Unit tests, mocks, benchmarks, and coverage reports

## 📋 Prerequisites

- Go 1.24+ 
- Slack Bot Token with appropriate permissions
- Channel IDs for the channels you want to compare

## 🛠️ Installation & Setup

**📖 For detailed setup instructions, see [docs/SETUP.md](docs/SETUP.md)**

### Quick Setup

1. **Clone the repository:**
```bash
git clone <repository-url>
cd slackVenn
```

2. **Install Go dependencies:**
```bash
go mod tidy
```

3. **Configure environment:**
```bash
# Copy template and edit with your token
cp env/.env.example env/.env
nano env/.env  # Add your SLACK_TOKEN
```

4. **Test installation:**
```bash
./slackVenn.sh --dry-run C1234567890 C0987654321
```

## 🔑 Getting a Slack Token

You need a Slack Bot Token with these scopes:
- `channels:read` - List public channels
- `groups:read` - List private channels  
- `users:read` - Get user information

**📖 Complete token setup guide: [docs/SETUP.md#getting-a-slack-bot-token](docs/SETUP.md#getting-a-slack-bot-token)**

**🔐 Private Channels:** For private channels, invite the app: `/invite @slackVenn Channel Analyzer`

## 🎯 Usage

### Basic Usage
```bash
# Automatic environment loading
./slackVenn.sh C1234567890 C0987654321

# Interactive mode (prompts for IDs)
./slackVenn.sh

# Dry run (test without creating files)
./slackVenn.sh --dry-run C1234567890 C0987654321
```

### Manual Environment Loading
```bash
# Load environment manually
source scripts/load-env.sh

# Then use any method
go run main.go C1234567890 C0987654321
./slackVenn C1234567890 C0987654321  # if compiled
```

### Finding Channel IDs

**From Slack URL:**
`https://yourworkspace.slack.com/archives/C1234567890` → Channel ID: `C1234567890`

**Using API:**
```bash
source scripts/load-env.sh
curl -H "Authorization: Bearer $SLACK_TOKEN" \
  "https://slack.com/api/conversations.list" | jq '.channels[] | {name: .name, id: .id}'
```

## 📊 Output Format

### Console Output
```bash
📊 slackVenn: Analyzing channel membership overlap...
🔍 Channel A: C1234567890
🔍 Channel B: C0987654321

📈 Analysis Results:
   Channel A: 45 members
   Channel B: 38 members
   Overlap: 12 members

🟢 Users in BOTH channels:
 - alice.johnson
 - bob.smith

🔵 Users ONLY in Channel A:
 - david.wilson

🟣 Users ONLY in Channel B:
 - frank.miller
```

### CSV Output
```csv
Status,Username
In Both,alice.johnson
In Both,bob.smith
Only in Channel A,david.wilson
Only in Channel B,frank.miller
```

## 🧪 Testing

slackVenn includes a comprehensive test suite with unit tests, integration tests, benchmarks, and coverage reports.

### Quick Test Run
```bash
# Run all tests with coverage and benchmarks
./scripts/run-tests.sh

# Run specific test types
go test ./tests/                    # Unit tests only
go test -bench=. ./tests/          # Benchmarks only
go test -v -run TestMock ./tests/  # Mock tests only
```

### Test Features
- ✅ **Unit Tests** - Core function testing with edge cases
- ✅ **Mock Tests** - Realistic scenarios without Slack API
- ✅ **Benchmarks** - Performance testing with large datasets
- ✅ **Coverage Reports** - HTML coverage reports with line-by-line analysis
- ✅ **Integration Tests** - End-to-end workflow validation

**📖 Complete testing guide: [tests/README.md](tests/README.md)**

## 📁 Project Structure

```
slackVenn/
├── env/                      # 🔐 Environment configuration
│   ├── .env.example         #    Template for environment variables
│   └── .env                 #    Your actual environment (gitignored)
├── scripts/                 # 🔧 Utility scripts
│   ├── load-env.sh         #    Environment loader script
│   └── run-tests.sh        #    Test suite runner
├── docs/                    # 📚 Documentation
│   └── SETUP.md            #    Complete setup guide
├── tests/                   # 🧪 Test suite
│   ├── main_test.go        #    Unit tests
│   ├── mock_test.go        #    Mock tests with realistic data
│   ├── results/            #    Test outputs and coverage reports
│   └── README.md           #    Testing documentation
├── history/                 # 📊 Generated comparison results
│   └── *.csv               #    Timestamped comparison files
├── main.go                  # 🚀 Core comparison logic
├── slackVenn.sh            # 🖥️  Shell script wrapper
├── go.mod                   # 📦 Go module definition
├── go.sum                   # 🔒 Go module checksums
├── .gitignore              # 🚫 Git ignore rules
├── CURRENT                 # 🔗 Symlink to latest result
└── README.md               # 📖 This file
```

## 🔧 Advanced Usage

### Batch Processing
```bash
# Compare multiple channel pairs
./slackVenn.sh C1111111111 C2222222222
./slackVenn.sh C3333333333 C4444444444
```

### Historical Analysis
```bash
# Track changes over time
./slackVenn.sh C1234567890 C0987654321

# Compare with previous results
diff history/channel_comparison_20250604_172759.csv \
     history/channel_comparison_20250611_172759.csv
```

### Building for Distribution
```bash
# Build standalone binary
go build -o slackVenn main.go

# Cross-platform builds
GOOS=linux GOARCH=amd64 go build -o slackVenn-linux main.go
GOOS=windows GOARCH=amd64 go build -o slackVenn.exe main.go
```

## 🎯 Use Cases

### Channel Auditing
- Find overlapping memberships across teams
- Identify users with access to sensitive channels
- Optimize channel structures for better organization

### Access Management
- Audit who has access to what channels
- Find users who should be in both channels but aren't
- Generate compliance reports

### Team Organization
- Understand team boundaries and overlaps
- Plan channel consolidation or splitting
- Analyze communication patterns

## 🔒 Security & Best Practices

- ✅ **Environment files are gitignored** - Your tokens stay private
- ✅ **Token validation** - Clear error messages for auth issues
- ✅ **Minimal permissions** - Only requests necessary Slack scopes
- ✅ **Local processing** - No data sent to external services
- ✅ **Test coverage** - Comprehensive testing ensures reliability

## 🐛 Troubleshooting

### Quick Fixes
```bash
# Environment issues
source scripts/load-env.sh

# Permission issues  
chmod +x slackVenn.sh scripts/load-env.sh scripts/run-tests.sh

# Dependency issues
go mod tidy

# Test your setup
./scripts/run-tests.sh
```

**📖 Complete troubleshooting guide: [docs/SETUP.md#troubleshooting](docs/SETUP.md#troubleshooting)**

## 🤝 Contributing & Sharing

### When sharing this project:
1. **Don't include your `.env` file** (already gitignored)
2. **Point users to [docs/SETUP.md](docs/SETUP.md)** for setup
3. **Each user needs their own Slack app/token**
4. **Run tests to verify everything works**: `./scripts/run-tests.sh`

### Ideas for enhancement:
- JSON output format for programmatic use
- Channel name resolution (display names, not just IDs)
- Web interface for easier channel selection
- Bulk comparison mode for multiple channel pairs

## 📄 License

MIT License - feel free to use, modify, and distribute.

---

**slackVenn** - Making Slack channel relationships visible, one Venn diagram at a time! 📊

**📖 Need help getting started?** Check out [docs/SETUP.md](docs/SETUP.md) for complete setup instructions. 