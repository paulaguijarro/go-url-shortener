package main

import (
	"github.com/paulaguijarro/go-url-shortener/database"
	"github.com/paulaguijarro/go-url-shortener/server"
)

func main() {
	database.ConnectDB()
	server.RunServer()
}
