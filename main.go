package main

import (
	"github.com/jSierraB3991/golang-multitenant/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	s := server.NewServer()
	s.Start()
}
