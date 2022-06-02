package socket

import (
	"interphlix/lib/socket/server"
	"log"
	"net"
)


var (
	Server server.Server
)


func Main() {
	go StartSocketServer()
	GetMyIPAddress()
}

func GetMyIPAddress() string {
	 conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr).IP.String()

    println(localAddr)
	return localAddr
}

func SetServer(IP string) {
	Servers, err := server.GetAllServers()
	HandlError(err)
	for _, server := range Servers {
		if server.LocalIPAddress == IP {
			Server = server
			return
		}
	}
	log.Panic("this server is not in the servers record")
}