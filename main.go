package main

import (
	"interphlix/lib/handler"
	"log"
	"net/http"
)

var (
	PORT = ":7000"
)

func main() {
	ConnectToDBs()
	go handler.Main()

	err := http.ListenAndServe(PORT, handler.Router)
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