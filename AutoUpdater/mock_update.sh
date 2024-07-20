#!/bin/bash

# Define the new version
NEW_VERSION="1.1.0"

# Update the version file
echo $NEW_VERSION > update_version.txt

# Compile the new binary with the version flag
echo "Compiling new version $NEW_VERSION..."
go build -ldflags "-X main.currentVersion=$NEW_VERSION" -o update main.go check_update_helper.go
chmod +x update
echo "New version $NEW_VERSION compiled as 'update'."
