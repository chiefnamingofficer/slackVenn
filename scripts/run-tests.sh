#!/bin/bash

# slackVenn Test Suite Runner
# Runs unit tests, integration tests, and benchmarks

set -e

echo "🧪 slackVenn Test Suite"
echo "======================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}📋 $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if we're in the right directory
if [ ! -f "main.go" ]; then
    print_error "run-tests.sh must be run from the slackVenn root directory"
    exit 1
fi

# Create test output directory
mkdir -p tests/results

# Generate timestamp for this test run
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
print_status "Test run timestamp: $TIMESTAMP"

print_status "Running unit tests..."

# Run unit tests with coverage
go test -v -cover -coverprofile=tests/results/coverage_${TIMESTAMP}.out ./tests/ 2>&1 | tee tests/results/test-output_${TIMESTAMP}.txt

# Check if tests passed
if [ ${PIPESTATUS[0]} -eq 0 ]; then
    print_success "Unit tests passed!"
else
    print_error "Unit tests failed!"
    exit 1
fi

print_status "Generating coverage report..."

# Generate HTML coverage report
go tool cover -html=tests/results/coverage_${TIMESTAMP}.out -o tests/results/coverage_${TIMESTAMP}.html

# Show coverage summary
COVERAGE=$(go tool cover -func=tests/results/coverage_${TIMESTAMP}.out | grep total | awk '{print $3}')
print_success "Test coverage: $COVERAGE"

print_status "Running benchmarks..."

# Run benchmarks
go test -bench=. -benchmem ./tests/ 2>&1 | tee tests/results/benchmark-output_${TIMESTAMP}.txt

print_status "Running performance tests..."

# Run performance tests
go test -v -run TestMockPerformance ./tests/ 2>&1 | tee tests/results/performance-output_${TIMESTAMP}.txt

print_status "Validating project structure..."

# Check that all expected files exist
EXPECTED_FILES=(
    "main.go"
    "slackVenn.sh"
    "go.mod"
    "README.md"
    "docs/SETUP.md"
    "env/.env.example"
    "scripts/load-env.sh"
)

for file in "${EXPECTED_FILES[@]}"; do
    if [ -f "$file" ]; then
        echo "✓ $file"
    else
        print_warning "Missing file: $file"
    fi
done

print_status "Testing shell script syntax..."

# Test shell script syntax
bash -n slackVenn.sh && echo "✓ slackVenn.sh syntax OK"
bash -n scripts/load-env.sh && echo "✓ load-env.sh syntax OK"

print_status "Testing dry-run functionality..."

# Test dry-run mode without requiring Slack token
./slackVenn.sh --dry-run C1234567890 C0987654321 > tests/results/dry-run-output_${TIMESTAMP}.txt 2>&1

if grep -q "✅ Dry run complete" tests/results/dry-run-output_${TIMESTAMP}.txt; then
    print_success "Dry-run test passed!"
else
    print_error "Dry-run test failed!"
    exit 1
fi

print_status "Creating symlinks to latest results..."

# Create symlinks to the latest results (similar to CURRENT in main app)
ln -sf coverage_${TIMESTAMP}.html tests/results/latest-coverage.html
ln -sf coverage_${TIMESTAMP}.out tests/results/latest-coverage.out
ln -sf test-output_${TIMESTAMP}.txt tests/results/latest-test-output.txt
ln -sf benchmark-output_${TIMESTAMP}.txt tests/results/latest-benchmark-output.txt
ln -sf performance-output_${TIMESTAMP}.txt tests/results/latest-performance-output.txt
ln -sf dry-run-output_${TIMESTAMP}.txt tests/results/latest-dry-run-output.txt

print_status "Test Summary"
echo "============"
echo "📁 Test results saved to: tests/results/"
echo "📊 Coverage report: tests/results/coverage_${TIMESTAMP}.html"
echo "⚡ Benchmark results: tests/results/benchmark-output_${TIMESTAMP}.txt"
echo "🎯 Performance results: tests/results/performance-output_${TIMESTAMP}.txt"
echo "🔗 Latest results: tests/results/latest-* (symlinks)"
echo ""

print_success "All tests completed successfully!"
print_status "Coverage: $COVERAGE"

# Optional: Open coverage report in browser (macOS)
if command -v open &> /dev/null; then
    read -p "Open coverage report in browser? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        open tests/results/coverage_${TIMESTAMP}.html
    fi
fi 