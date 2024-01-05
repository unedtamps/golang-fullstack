package main

import "github.com/unedtamps/golang-fullstack/bootstrap"

func main() {
	server := bootstrap.LoadServer()
	server.Start()
}

