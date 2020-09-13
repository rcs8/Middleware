package socket

import (
	"encoding/json"
	"fmt"
	"math"
	"net"

	"../shared"
)

type ServerUDP struct {
	conn *net.UDPConn
}

func NewServerUDP(port string) (*ServerUDP, error) {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return nil, err
	}

	// Listen on udp port
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &ServerUDP{
		conn: conn,
	}, err
}

func (s *ServerUDP) ListenUDP() {
	msgFromClient := make([]byte, 1024)
	for {
		// Receive request
		n, addr, err := (*s.conn).ReadFromUDP(msgFromClient)
		if err != nil {
			fmt.Println(err)
		}

		// Handle request
		go HandleUDP(s.conn, msgFromClient, n, addr)
	}
}

func HandleUDP(conn *net.UDPConn, msgFromClient []byte, n int, addr *net.UDPAddr) {
	var msgToClient []byte
	var args shared.Args

	err := json.Unmarshal(msgFromClient[:n], &args)
	if err != nil {
		fmt.Println(err)
	}

	result := InvokeSqrt(args)

	msgToClient, err = json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.WriteTo(msgToClient, addr)
	if err != nil {
		fmt.Println(err)
	}
}

func InvokeSqrt(args shared.Args) shared.Reply {
	var a = float64(args.A)
	var b = float64(args.B)
	var c = float64(args.C)

	deltaValue := CalculateDelta(a, b, c)

	if deltaValue < 0 {
		return shared.Reply{
			Result: "Nenhuma raiz real\n",
		}
	}

	if deltaValue == 0 {
		return shared.Reply{
			Result: fmt.Sprintf("%f\n", (b*(-1))/(2*a)),
		}
	}

	return shared.Reply{
		Result: fmt.Sprintf("%f e %f\n", (math.Sqrt(deltaValue)-b)/2*a, ((-1)*math.Sqrt(deltaValue)-b)/2*a),
	}
}

func CalculateDelta(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}
