package main

import (
	"food_delivery/server"
	_ "github.com/lib/pq"
)

func main() {
	server.Start()
}
