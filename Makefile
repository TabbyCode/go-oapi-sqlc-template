.PHONY: all install clean format lint test build run

all: clean build

install:
	# Installing Moonrepo proto https://moonrepo.dev/proto
	bash -c "$$(curl -fsSL https://moonrepo.dev/install/proto.sh)"
	# Installing all the other tools with the proto
	proto use
	# Running the Moon setup task at the end
	moon :setup

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
