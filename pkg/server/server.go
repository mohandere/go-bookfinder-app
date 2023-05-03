package server

import "go-playground/pkg/db"

func Start() {
	r := NewRouter()
	db.Init()
	r.Run()
}