package main

import (
	"fmt"
	"go-bookfinder-app/pkg/server"
)

func main() {
	fmt.Println("Starting web server...")
	server.Start()
}