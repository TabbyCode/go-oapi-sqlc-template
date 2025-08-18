#!/usr/bin/env bash

# Go files
golangci-lint run -c configs/golangci.yml --fix --allow-parallel-runners --issues-exit-code 0 > /dev/null

# SQL files
sqruff fix --format github-annotation-native --config configs/sqruff.toml db > /dev/null 2> /dev/null
