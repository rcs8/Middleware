package main

import (
	"fmt"
	"net"
)

func main() {
	srhTCP, err := newSRHTCP()
	if err != nil {
		panic(err)
	}
	srhTCP.Receive()
}

func responseMessage() []byte {
	var message = "this is what server said."
	return []byte(message)
}

func HandleTCP(srh srhTCP, conn net.Conn, message []byte) {

	defer conn.Close()
	s := string(message[:])
	fmt.Print(string(s), "from client\n")
	var msgToClient []byte = responseMessage()

	srh.SendTCP(msgToClient, conn)
}
