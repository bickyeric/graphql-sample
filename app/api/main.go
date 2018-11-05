package main

import (
	"github.com/bickyeric/scaling-chainsaw/server"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")
	server.Start()
}
