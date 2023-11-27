# Library App 

This project contains two different applications. The Library App Server and the Library App CLI. The server is a REST API that allows users to create, delete, update, get books and book collections. The CLI is a command line interface that allows users to interact with the server. 

## Server

The server uses SQLite to store the book and book collection data as well as [gorm](./gorm.db) to interact with the database. The REST multiplexed used is [gorilla/mux](https://github.com/gorilla/mux). 

To run the server, run:

``` bash
cd server
go run .
```

The server has the option to supply a port number to listen on as well as a place to load and store the SQLite database file. The default port is 8080 and the default database file is `server/gorm.db`.

As an example, to run the server with either of these paramter's changed, execute: 

``` bash
cd server
PORT=8000 DB_FILE=/tmp/libraryapp/gorm.db go run . --port=8081 --dbfile=/tmp/gorm.db
```

##

- [ ] Add tests
- [ ] Add authentication
- [ ] Add rate limiting

## CLI

The CLI uses [cobra]()