# aero

A simple `reverse proxy` CLI tool made in [Go](https://go.dev/).

Created to practice (and learn) about proxies and reverse proxies.

![aero](https://frutigeraeroarchive.org/images/introduction/asadal.jpg)

## Project Structure

```
aero/
├── cmd/
|   ├── server/
|       ├── go.mod
|       ├── main.go  # Basic go HTTP server
|   ├── root.go      # Base command
|   ├── run.go       # Main subcommand
├── internal/        # App internal logic
|   ├── app/
|   ├── balancer/
|   ├── config/
|   ├── logger/
├── config.yaml      # Configuration file
├── .gitignore
├── go.mod
├── main.go
```

## Prerequisites

1. Golang 1.21 or higher ([install here](https://go.dev/dl/))!
2. Git installed on your machine.

## Installation

1. Clone the repo on your machine.
2. Run `go install .`
3. Open one (or more) terminals separated from the main one. In each terminal, start a simple Go server with:
```
cd .\cmd\server
go run main.go
```

4. Start the main app:
```
go run main.go
```

You can also install it in your machine with:
```
go build; go install
```

And start it:
```
aero run
```

5. If you want, you can use the `--verbose` flag to output more details about the operations of the app.
```
aero run --verbose
```

## Thanks for visiting!
