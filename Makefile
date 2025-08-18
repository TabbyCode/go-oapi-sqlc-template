.PHONY: all install clean setup generate format lint test build serve

all: clean build

install:
	bash scripts/install.sh

clean:
	moon :clean

setup:
	moon :setup

generate:
	moon :generate

format:
	moon :format

lint:
	moon :lint

test:
	moon :test

build:
	moon :build

serve:
	moon :serve

migrate:
	moon :migrate
