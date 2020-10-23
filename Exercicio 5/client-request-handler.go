package main

import {

}

type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
}

func (crh ClientRequestHandler) SendReceive(msgToServer []byte) []byte {
	
} 