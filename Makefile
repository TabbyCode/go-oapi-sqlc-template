.PHONY: all install clean format lint test build run

all: clean build

install:
	bash scripts/install.sh

clean:
	moon :clean

format:
	moon :format

lint:
	moon :lint

test:
	moon :test

build:
	moon :build

run:
	moon :run
