package main

import (
	"interphlix/lib/handler"
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

	err := http.ListenAndServe(HTTPPORT, handler.Router)
	HandlError(err)
}

/// connect to both the local and remote mongodb databases
func ConnectToDBs() {
	ConnectToRemoteDB1()
	ConnectRemotedDB2()
	ConnectLocalDB()
}


// handle errors that need the program to exit
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}