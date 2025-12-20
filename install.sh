#!/bin/sh

set -e

# --- Build and Install Binary ---

INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

echo "Building mm..."
go build -o mm ./cmd/makemego

echo "Installing mm to $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
if [ -w "$INSTALL_DIR" ]; then
    mv mm "$INSTALL_DIR/mm"
else
    sudo mv mm "$INSTALL_DIR/mm"
fi

echo "mm binary installed successfully!"
echo

# --- Shell Integration ---

# Ask the user if they want to install shell integration
# If NONINTERACTIVE is set, assume yes
if [ -z "$NONINTERACTIVE" ]; then
    printf "Install shell integration? (y/N) "
    read -r answer
    if [ "$answer" != "y" ]; then
        echo "Skipping shell integration."
        exit 0
    fi
fi

# Detect shell
if [ -n "$PROFILE_FILE" ]; then
    # If PROFILE_FILE is provided, we need to know the shell type to generate correct function.
    # We'll assume bash if not specified via SHELL_TYPE env var, or try to detect.
    if [ -z "$SHELL_TYPE" ]; then
        PARENT_SHELL=$(ps -o comm= -p $PPID 2>/dev/null || echo "")
        PARENT_SHELL_NAME=$(basename "$PARENT_SHELL" | sed 's/-//')

        if echo "$PARENT_SHELL_NAME" | grep -q "zsh"; then
             SHELL_TYPE="zsh"
        elif echo "$PARENT_SHELL_NAME" | grep -q "bash"; then
             SHELL_TYPE="bash"
        elif echo "$SHELL" | grep -q "zsh"; then
             SHELL_TYPE="zsh"
        elif [ -n "$BASH_VERSION" ]; then
            SHELL_TYPE="bash"
        elif [ -n "$ZSH_VERSION" ]; then
            SHELL_TYPE="zsh"
        else
            SHELL_TYPE="bash" # Default to bash for testing if not detected
        fi
    fi
elif ps -o comm= -p $PPID 2>/dev/null | grep -q "zsh"; then
    SHELL_TYPE="zsh"
    PROFILE_FILE="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
    SHELL_TYPE="bash"
    PROFILE_FILE="$HOME/.bashrc"
elif [ -n "$ZSH_VERSION" ]; then
    SHELL_TYPE="zsh"
    PROFILE_FILE="$HOME/.zshrc"
else
    echo "Unsupported shell for automatic integration. Please set up the 'mm' function manually."
    exit 0
fi

echo "Detected $SHELL_TYPE shell. Will add 'mm' function to $PROFILE_FILE."

# Create the function string
if [ "$SHELL_TYPE" = "bash" ]; then
    FUNCTION_STRING='''
# MakeMeGo shell integration
mm() {
  local selected_command
  selected_command=$(command mm --print-command "$@")
  if [ -n "$selected_command" ]; then
    history -s "$selected_command"
    eval "$selected_command"
  fi
}
'''
elif [ "$SHELL_TYPE" = "zsh" ]; then
    FUNCTION_STRING='''
# MakeMeGo shell integration
mm() {
  local selected_command
  selected_command=$(command mm --print-command "$@")
  if [ -n "$selected_command" ]; then
    print -z "$selected_command"
  fi
}
'''
fi

# Append to profile
echo "$FUNCTION_STRING" >> "$PROFILE_FILE"

echo
echo "'mm' function added to $PROFILE_FILE."
echo "Please restart your shell or run 'source $PROFILE_FILE' to apply the changes."