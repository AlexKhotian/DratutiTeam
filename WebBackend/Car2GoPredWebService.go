package main

import "DratutiTeam/WebBackend/HTTPHandler"

func main() {
	server := HTTPHandler.ServerRoutineFactory()
	server.RunServer()
}