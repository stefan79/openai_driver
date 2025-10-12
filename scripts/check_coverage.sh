#!/bin/bash

# Exit on any error
set -e

# Define minimum coverage requirements for packages
# Using a space-separated list of package=coverage pairs
MIN_COVERAGE=(
    "driver/pkg/builder=60"
    "driver/pkg/cli=60"
    # Add more packages and their minimum coverage requirements here
)

# Get coverage for all packages
coverage_data=$(go test -coverprofile=coverage.txt -covermode=atomic ./... | grep -E '^ok\s+[^\s]+\s+.*coverage: [0-9.]+%' || true)

if [ -z "$coverage_data" ]; then
    echo "No test coverage data found. Make sure tests are running correctly."
    exit 1
fi

echo "Checking package coverage..."
# Process each line of coverage data
while IFS= read -r line; do
    # Extract package name and coverage percentage
    pkg=$(echo "$line" | grep -o '^ok[[:space:]]\+[^[:space:]]\+' | awk '{print $2}')
    coverage=$(echo "$line" | grep -o 'coverage: [0-9.]\+%' | cut -d' ' -f2 | tr -d '%')
    
    # Skip if we couldn't parse the package or coverage
    if [ -z "$pkg" ] || [ -z "$coverage" ]; then
        continue
    fi
    
    # Check if this package has a minimum coverage requirement
    for item in "${MIN_COVERAGE[@]}"; do
        pkg_pattern="${item%=*}"  # Get the part before =
        min_coverage="${item#*=}"  # Get the part after =
        
        if [[ "$pkg" == *"$pkg_pattern"* ]]; then
            # Compare coverage with minimum required
            if (( $(echo "$coverage < $min_coverage" | bc -l 2>/dev/null) )); then
                echo "❌ Package $pkg has ${coverage}% test coverage, which is below the required ${min_coverage}%"
                exit 1
            else
                echo "✅ Package $pkg has ${coverage}% test coverage (minimum required: ${min_coverage}%)"
            fi
            break
        fi
    done
done <<< "$coverage_data"

echo "All packages meet or exceed their minimum coverage requirements!"
exit 0
