// package main

// import (
// 	"encoding/json"
// 	"math/rand"
// 	"net"
// 	"time"
// )

// type Client struct {
// 	conn    *net.Conn
// 	encoder *json.Encoder
// 	decoder *json.Decoder
// }

// func NewClientTCP(address string) (*Client, error) {
// 	conn, err := net.Dial("tcp", address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return CreateEncoders(conn), err
// }

// func NewClientUDP(address string) (*Client, error) {
// 	addr, err := net.ResolveUDPAddr("udp", address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conn, err := net.DialUDP("udp", nil, addr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return CreateEncoders(conn), err
// }

// func CreateEncoders(conn net.Conn) *Client {
// 	jsonEncoder := json.NewEncoder(conn)
// 	jsonDecoder := json.NewDecoder(conn)

// 	return &Client{
// 		conn:    &conn,
// 		encoder: jsonEncoder,
// 		decoder: jsonDecoder,
// 	}
// }

// func (c *Client) MakeRequest() ([]float64, error) {
// 	var response Reply
// 	var err error

// 	message := PrepareArgs()

// 	err = c.encoder.Encode(message)

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = c.decoder.Decode(&response)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response.Result, err
// }

// func (c *Client) MakeRequestBenchmark() ([]float64, int64, error) {
// 	var response Reply
// 	var err error

// 	message := PrepareArgs()

// 	startTime := time.Now()
// 	err = c.encoder.Encode(message)

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	err = c.decoder.Decode(&response)
// 	totalTime := time.Now().Sub(startTime).Microseconds()

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return response.Result, totalTime, err
// }

// func PrepareArgs() Args {
// 	rand.Seed(time.Now().UnixNano())
// 	return Args{
// 		A: rand.Intn(10) + 1,
// 		B: rand.Intn(10) + 1,
// 		C: rand.Intn(10) + 1,
// 	}
// }

// func (c *Client) Close() {
// 	(*c.conn).Close()
// }

package main

import (
  "log"
  "gitlab.com/pantomath-io/demo-grpc/api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)
func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":6666", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  c := api.NewSqrtServiceClient(conn)
  response, err := c.Sqrt(context.Background(), &api.Args{A: 1, B: 2, C: 3})
  if err != nil {
    log.Fatalf("Error when calling Sqrt: %s", err)
  }
  log.Printf("Response from server: %s", response.Result)
}