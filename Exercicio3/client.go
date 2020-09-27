package main

import (
  "log"
  
  "ex3/api"
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
  args := api.Args{
    A: 1,
    B: -3,
    C: -10,
  }
  response, err := c.Sqrt(context.Background(), &args)
  if err != nil {
    log.Fatalf("Error when calling Sqrt: %s", err)
  }
  log.Printf("Response from server: %s", response.Result)
}