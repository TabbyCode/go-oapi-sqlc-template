#!/usr/bin/env bash

mkdir -p internal/gen/db internal/gen/oapi pkg/gen

oapi-codegen -config oapi-codegen-server.yml ./api/openapi.yml
oapi-codegen -config oapi-codegen-client.yml ./api/openapi.yml

sqlc generate -f sqlc.yml
