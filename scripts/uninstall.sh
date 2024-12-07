#!/bin/bash

# uninstall.sh

INSTALL_DIR="/usr/local/bin"
BINARY_NAME="go-start"

if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
    echo "Uninstalling application..."
    sudo rm "$INSTALL_DIR/$BINARY_NAME"
    echo "Uninstallation complete"
else
    echo "Application is not installed"
fi
