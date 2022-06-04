package main

import "merchant-service/server"

func main() {
	merchantServer := server.NewServer()
	merchantServer.Start()
}
