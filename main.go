package main

import (
	b "github.com/unedtamps/chat-app/bootstrap"
)

func main(){
	server := b.LoadServer()
	server.Start()
}