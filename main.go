package main

import (
	"github.com/DuduNeves/Estoque_v1/database"
	"github.com/DuduNeves/Estoque_v1/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()
}
