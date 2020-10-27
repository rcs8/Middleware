package main

import (
	"fmt"
	"net"
)

func main() {
	srhUDP, err := newSRHUDP()
	if err != nil {
		panic(err)
	}
	srhUDP.Receive()
}

func responseMessage() []byte {
	var message = "this is what server said."
	return []byte(message)
}

func HandleUDP(srh srhUDP, conn net.UDPConn, message []byte, addr *net.UDPAddr) {
	s := string(message[:])
	fmt.Print(string(s), "from client\n")
	var msgToClient []byte = responseMessage()

	srh.SendUDP(msgToClient, conn, addr)
}
