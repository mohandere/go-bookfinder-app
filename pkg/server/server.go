package server

import "go-bookfinder-app/pkg/db"

func Start() {
	r := NewRouter()
	db.Init()
	db.PrintDBConn()
	r.Run()
}