package socket

import (
	"log"
	"net"
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