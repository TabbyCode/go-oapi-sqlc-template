#!/usr/bin/env bash

mkdir -p internal/gen/db internal/gen/oapi pkg/gen

go tool oapi-codegen -config ./configs/oapi-codegen-server.yml ./api/openapi.yml
go tool oapi-codegen -config ./configs/oapi-codegen-client.yml ./api/openapi.yml

sqlc generate -f ./configs/sqlc.yml
