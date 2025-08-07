#!/bin/bash

# Create directories if they don't exist
mkdir -p internal/gen pkg/gen

# Generate server code
oapi-codegen -config oapi-codegen-server.yml ./api/openapi.yml

# Generate client code
oapi-codegen -config oapi-codegen-client.yml ./api/openapi.yml
