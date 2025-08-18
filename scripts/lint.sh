#!/usr/bin/env bash

# Go files
golangci-lint run -c configs/golangci.yml --allow-parallel-runners

# SQL files
sqruff lint --format github-annotation-native --config configs/sqruff.toml db
