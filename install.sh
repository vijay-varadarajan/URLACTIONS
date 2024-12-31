#!/bin/bash

# Determine the OS
OS=$(uname -s)
if [[ "$OS" == "Linux" ]]; then
    BINARY_NAME="urlactions"
elif [[ "$OS" == "Darwin" ]]; then
    BINARY_NAME="urlcli"
else
    echo "Unsupported operating system: $OS"
    exit 1
fi

DESTINATION="/usr/local/bin"

# Ensure the script is run with root privileges
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root. Try 'sudo ./install.sh'."
   exit 1
fi

# Check if the binary exists in the current directory
if [[ ! -f "$BINARY_NAME" ]]; then
    echo "Error: $BINARY_NAME not found in the current directory."
    exit 1
fi

# Make the binary executable
chmod +x "$BINARY_NAME"

# Move the binary to the destination directory
echo "Installing $BINARY_NAME to $DESTINATION..."
mv "$BINARY_NAME" "$DESTINATION"

# Verify installation
if [[ $? -eq 0 ]]; then
    echo "Installation successful!"
    echo "You can now run the tool using: $BINARY_NAME"
else
    echo "Installation failed."
    exit 1
fi
