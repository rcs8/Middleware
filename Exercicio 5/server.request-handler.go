package main

import {
	"net"
}

type ServerRequestHandler struct {}

var ln net.Listener
var conn net.Conn
var err error

func (srh ServerRequestHandler) Receive() []byte {

}

func (srh ServerRequestHandler) Send(msgToClient []byte) {

}