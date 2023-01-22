#!/bin/bash

# Build the Go binary
GOOS=linux GOARCH=amd64 go build -o main

# Get external dependencies
go get -d -v ./...

# Create the deployment package
zip main.zip main

# Deploy the function using Terraform
terraform init
terraform apply -auto-approve

# Clean up
rm main
rm main.zip

