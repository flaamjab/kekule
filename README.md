# Kekule

A very basic RESTful API for a fictional online shop.

## Compiling and Running

The API uses go-sqlite3 and therefore must be compiled with `CGO_ENABLED=1`
environment variable; gcc must also be present in `PATH`.

To start the API server use the following command:

    go run cmd/kekule/server.go
