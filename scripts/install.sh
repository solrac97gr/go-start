#!/bin/bash

# install.sh

# Colors for messages
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed${NC}"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

# Default installation directory
INSTALL_DIR="/usr/local/bin"
# Binary name (change according to your project)
BINARY_NAME="go-start"

echo -e "${GREEN}Starting installation...${NC}"

# Build the project
echo "Building project..."
if ! go build -o "$BINARY_NAME" main.go; then
    echo -e "${RED}Error: Build failed${NC}"
    exit 1
fi

# Create installation directory if it doesn't exist
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating installation directory..."
    sudo mkdir -p "$INSTALL_DIR"
fi

# Move binary to installation directory
echo "Installing binary..."
if sudo mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"; then
    echo -e "${GREEN}Installation completed successfully${NC}"
    echo "Binary available at: $INSTALL_DIR/$BINARY_NAME"
else
    echo -e "${RED}Error: Could not install binary${NC}"
    exit 1
fi

# Verify installation
if command -v "$BINARY_NAME" &> /dev/null; then
    echo -e "${GREEN}Application installed successfully and available in PATH${NC}"
else
    echo -e "${RED}Warning: Application installed but might not be in PATH${NC}"
fi

echo -e "${GREEN}Installation complete!${NC} 🎊"
