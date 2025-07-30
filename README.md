# The Awesome Go project

- This project adopts the [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for a clean
  and organized codebase.

- We recommend following the guidelines outlined in [Effective Go](https://go.dev/doc/effective_go). Most formatting can
  be automatically handled using the `go fmt` tool.

- For projects involving API design, we highly recommend referring to the
  [Google API Design Guide](https://cloud.google.com/apis/design) for best practices.

## Setup

We can install dependencies with:

```shell
go mod download
```

also, it's possible to use the `make`:

```shell
make install
```

and to run the code:

```shell
go cmd/main.go
```

## Build

To simply build the project a simple `make` command is enough, and then you could find the binary file in the `bin`
directory:

```shell
make
```

to see all the available commands, use `help` argument:

```shell
make help
```
