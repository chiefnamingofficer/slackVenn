# slackVenn Test Suite

Comprehensive testing framework for the slackVenn project including unit tests, integration tests, benchmarks, and mock testing.

## 🚀 Quick Test Run

```bash
# Run all tests with timestamped results
./scripts/run-tests.sh

# Run specific test types
go test ./tests/                    # Unit tests only
go test -bench=. ./tests/          # Benchmarks only
go test -v -run TestMock ./tests/  # Mock tests only
```

## 📁 Test Structure

```
tests/
├── main_test.go         # Unit tests for core functions
├── mock_test.go         # Mock Slack API tests
├── results/             # Test output and reports (gitignored)
│   ├── .gitkeep        #   Preserves directory structure in git
│   ├── *_YYYYMMDD_HHMMSS.html   #   Timestamped HTML coverage reports
│   ├── *_YYYYMMDD_HHMMSS.out    #   Timestamped coverage data
│   ├── *_YYYYMMDD_HHMMSS.txt    #   Timestamped test outputs
│   └── latest-*        #   Symlinks to most recent results
└── README.md           # This file
```

## 🔒 Git Ignore Behavior

**Files ignored by git:**
- ✅ `tests/results/*` - All test output files (except `.gitkeep`)
- ✅ `history/*.csv` - Channel comparison results (user data)
- ✅ `CURRENT` - Symlink to latest comparison result

**Files tracked by git:**
- ✅ `tests/results/.gitkeep` - Preserves directory structure
- ✅ `history/.gitkeep` - Preserves directory structure
- ✅ All source code and documentation

This ensures:
- **No sensitive data** in git (user lists from Slack)
- **No generated files** cluttering the repository
- **Clean directory structure** preserved for new checkouts

## 🧪 Test Categories

### 1. Unit Tests (`main_test.go`)

Tests core functions with various input scenarios:

- **`TestDifference`** - Tests the difference function with edge cases
- **`TestIntersection`** - Tests the intersection function with edge cases
- **`BenchmarkDifference`** - Performance benchmarks for difference operations
- **`BenchmarkIntersection`** - Performance benchmarks for intersection operations

**Run unit tests:**
```bash
go test -v ./tests/ -run "Test(Difference|Intersection)"
```

### 2. Mock Tests (`mock_test.go`)

Tests application logic with mock Slack data:

- **`TestMockChannelComparison`** - Tests realistic channel comparison scenarios
- **`TestMockEdgeCases`** - Tests edge cases like empty channels, large datasets
- **`TestMockPerformance`** - Tests performance with large mock datasets

**Run mock tests:**
```bash
go test -v ./tests/ -run "TestMock"
```

### 3. Integration Tests (via `run-tests.sh`)

Tests the complete application flow:

- **Shell script syntax validation**
- **Dry-run functionality testing**
- **Project structure validation**
- **Environment loading testing**

## 📊 Coverage Reports

The test suite generates detailed coverage reports with timestamps:

```bash
# Generate coverage report with timestamp
./scripts/run-tests.sh

# View latest coverage in browser (macOS)
open tests/results/latest-coverage.html

# View specific timestamped coverage
open tests/results/coverage_20250604_182617.html

# View coverage summary
go tool cover -func=tests/results/latest-coverage.out
```

## ⚡ Performance Testing

### Benchmarks

Test performance of core algorithms:

```bash
# Run all benchmarks
go test -bench=. ./tests/

# Run specific benchmarks
go test -bench=BenchmarkDifference ./tests/
go test -bench=BenchmarkIntersection ./tests/

# Benchmark with memory allocation stats
go test -bench=. -benchmem ./tests/
```

### Large Dataset Testing

Tests with realistic large channel sizes:

```bash
# Test with large mock datasets
go test -v -run TestMockPerformance ./tests/
```

## 🎯 Mock Data

The test suite includes realistic mock data:

- **5 Mock Users**: alice.johnson, bob.smith, charlie.brown, david.wilson, emma.davis
- **2 Mock Channels**: 
  - `C1234567890`: 4 members (alice, bob, charlie, david)
  - `C0987654321`: 3 members (alice, bob, emma)
- **Expected Results**:
  - Common: 2 users (alice, bob)
  - Only in A: 2 users (charlie, david) 
  - Only in B: 1 user (emma)

## 🔧 Testing Without Slack Token

All tests can run without a real Slack token:

- **Unit tests** use pure Go functions
- **Mock tests** use predefined test data
- **Integration tests** use dry-run mode

This enables:
- ✅ Testing in CI/CD pipelines
- ✅ Testing without Slack workspace access
- ✅ Consistent test results
- ✅ Fast test execution

## 📈 Test Metrics

The test suite tracks with timestamped files:

- **Code Coverage** - Percentage of code tested (`coverage_TIMESTAMP.html`)
- **Performance** - Function execution times (`benchmark-output_TIMESTAMP.txt`)
- **Memory Usage** - Memory allocation patterns (in benchmark files)
- **Edge Cases** - Boundary condition handling (in test outputs)
- **Historical Tracking** - Multiple test runs preserved

## 🐛 Debugging Failed Tests

### Unit Test Failures

```bash
# Run with verbose output
go test -v ./tests/ -run TestDifference

# Run specific test case
go test -v ./tests/ -run "TestDifference/Simple_difference"
```

### Mock Test Failures

```bash
# Run with detailed logging
go test -v ./tests/ -run TestMockChannelComparison
```

### Integration Test Failures

Check specific timestamped test outputs:

```bash
# Check latest dry-run output
cat tests/results/latest-dry-run-output.txt

# Check latest full test log
cat tests/results/latest-test-output.txt

# Check specific timestamped results
cat tests/results/dry-run-output_20250604_182617.txt
```

## 🚀 Adding New Tests

### Adding Unit Tests

1. Add test function to `main_test.go`:
```go
func TestYourFunction(t *testing.T) {
    // Test implementation
}
```

2. Run specific test:
```bash
go test -v ./tests/ -run TestYourFunction
```

### Adding Mock Tests

1. Add mock data to `mock_test.go`
2. Create test function with realistic scenarios
3. Validate expected vs actual results

### Adding Integration Tests

1. Add test steps to `scripts/run-tests.sh`
2. Include validation logic
3. Update expected files list if needed

## 📄 Test Best Practices

- ✅ **Test edge cases** - Empty inputs, large datasets, boundary conditions
- ✅ **Use descriptive names** - Clear test function and case names
- ✅ **Include benchmarks** - Performance regression detection
- ✅ **Mock external APIs** - No external dependencies in tests
- ✅ **Validate error handling** - Test failure scenarios
- ✅ **Generate coverage reports** - Track test completeness

---

**Run the full test suite anytime with:** `./scripts/run-tests.sh` 