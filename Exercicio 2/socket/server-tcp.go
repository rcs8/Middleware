package socket

import (
	"encoding/json"
	"fmt"
	"math"
	"net"

	"../shared"
)

type ServerTCP struct {
	listener *net.Listener
}

func NewServerTCP(port string) (*ServerTCP, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &ServerTCP{
		listener: &ln,
	}, err
}

func (s *ServerTCP) ListenTCP() {
	for {
		conn, err := (*s.listener).Accept()
		if err != nil {
			fmt.Println(err)
		}

		go HandleConnectionTCP(conn)
	}
}

func HandleConnectionTCP(conn net.Conn) {
	var messageFromClient shared.Args

	defer conn.Close()

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		err := jsonDecoder.Decode(&messageFromClient)
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}

		msgToClient := InvokeSqrt(messageFromClient)

		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
		}
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
