package main

import (
	"interphlix/lib/handler"
	"interphlix/lib/socket"
	"interphlix/lib/variables"
	"log"
	"net/http"
)

var (
	HTTPPORT = ":7000"
	SOCKETPORT = ":9000"
)

func main() {
	ConnectToDBs()
	go handler.Main()
	go socket.Main()
	println("starting server...")

	err := http.ListenAndServe(HTTPPORT, handler.Router)
	HandlError(err)
}

/// connect to both the local and remote mongodb databases
func ConnectToDBs() {
	variables.ConnectToRemoteDB1()
	variables.ConnectRemotedDB2()
	variables.ConnectLocalDB()
}


// handle errors that need the program to exit
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}