# Library App 

This is a portfolio project that contains two different applications: The Library App Server and the Library App CLI. The server is a REST API that allows users to create, delete, update, and load books as well as book collections. The CLI is designed to allow users to interact with the API.

## Getting Started

The server uses SQLite to store the book and book collection data as well as [gorm](./gorm.db) to interact with the database. The REST framework used is [gin-gonic/gin](http://github.com/gin-gonic/gin), and the CLI framework is [spf13/cobra]("github.com/spf13/cobra"). To install all dependencies run `go get` from the root of the project

## Server

To run the server, run:

``` bash
export GIN_MODE=release go run server/main.go
```

The server has the option to supply a port number to listen on as well as a place to load and store the SQLite database file. The default port is 8080 and the default database file is `server/gorm.db`. As an example, to run the server with either of these paramter's changed, execute: 

``` bash
PORT=8000 GIN_MODE=release DB_FILE=/tmp/libraryapp/gorm.db go run server/main.go
```

To access the API documentation for the server, navigate to `http://localhost:8080/swagger/index.html` in your browser

## CLI

To run the cli tool, execute: 

`go run cli/main.go [command]`

## Additional Documents

1. [Expected User Experience](./docs/expected_user_experience.md)
1. [Implementation And Assumption Notes](./docs/implementation_and_assumption_notes.md)