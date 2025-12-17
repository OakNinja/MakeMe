#!/bin/bash

set -e

# Setup temp directory
TEST_DIR=$(mktemp -d)
INSTALL_DIR="$TEST_DIR/bin"
PROFILE_FILE="$TEST_DIR/.bashrc"
XDG_CONFIG_HOME="$TEST_DIR/.config"

cleanup() {
    rm -rf "$TEST_DIR"
}
trap cleanup EXIT

echo "Running tests in $TEST_DIR"

# --- Test install.sh (Bash) ---
echo "Testing install.sh (Bash)..."
export INSTALL_DIR
export PROFILE_FILE
export NONINTERACTIVE=1
export SHELL_TYPE=bash

# Create dummy profile
touch "$PROFILE_FILE"

# Run install.sh
./install.sh

# Verify binary
if [ ! -f "$INSTALL_DIR/mm" ]; then
    echo "FAIL: mm binary not found in $INSTALL_DIR"
    exit 1
fi

# Verify profile update
if ! grep -q "mm()" "$PROFILE_FILE"; then
    echo "FAIL: mm function not found in $PROFILE_FILE"
    exit 1
fi

echo "install.sh (Bash) passed!"

# --- Test install.fish ---
echo "Testing install.fish..."
# Reset binary to verify fish installs it too (or overwrites it)
rm "$INSTALL_DIR/mm"

# Fish uses XDG_CONFIG_HOME
export XDG_CONFIG_HOME

# Run install.fish
# We need fish installed to run this. If not available, skip.
if command -v fish >/dev/null; then
    fish ./install.fish
    
    # Verify binary
    if [ ! -f "$INSTALL_DIR/mm" ]; then
        echo "FAIL: mm binary not found in $INSTALL_DIR after fish install"
        exit 1
    fi

    # Verify fish function
    FISH_FUNC="$XDG_CONFIG_HOME/fish/functions/mm.fish"
    if [ ! -f "$FISH_FUNC" ]; then
        echo "FAIL: mm.fish not found at $FISH_FUNC"
        exit 1
    fi

    if ! grep -q "function mm" "$FISH_FUNC"; then
        echo "FAIL: 'function mm' not found in $FISH_FUNC"
        exit 1
    fi

    echo "install.fish passed!"
else
    echo "fish not found, skipping install.fish test."
fi

echo "All installation tests passed!"
