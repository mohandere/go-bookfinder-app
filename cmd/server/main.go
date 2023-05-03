package main

import (
	"fmt"
	"go-playground/pkg/server"
)

func main() {
	fmt.Println("Starting web server...")
	server.Start()
}