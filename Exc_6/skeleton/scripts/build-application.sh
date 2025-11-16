#!/bin/sh
# Exit immediately if any command fails
set -e

# Change to app directory
cd /app || exit 1

# Download Go modules
go mod download

# Build the Go binary for Linux
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem