package main
import (
	"log"
	"net"
	"fmt"
	"ex3/api"
  	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6666))
	if err != nil {
	  log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listenin port 6666")

	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterSqrtServiceServer(grpcServer, &s)

	log.Printf("grpcServer registered")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC.")
	}

	// api.RegisterSqrtServiceServer(grpcServer, &server{})
	// reflection.Register(grpcServer)
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }
}

// func (s *Server) Sqrt(ctx context.Context, args *api.Args) (*api.Reply, error) {
// 	log.Printf("Receive new message")
// 	result := []float64{}
// 	var a = args.GetA()
// 	var b = args.GetB()
// 	var c = args.GetC()

// 	deltaValue := CalculateDelta(a, b, c)

// 	if deltaValue < 0 {
// 		return &proto.Reply{
// 			Result: result,
// 		}, nil
// 	}

// 	if deltaValue == 0 {
// 		return &proto.Reply{
// 			Result: append(result, (b*(-1))/(2*a)),
// 		}, nil
// 	}

// 	return &proto.Reply{
// 		Result: append(result, (math.Sqrt(deltaValue)-b)/2*a, ((-1)*math.Sqrt(deltaValue)-b)/2*a),
// 	}, nil
// }

// func CalculateDelta(a, b, c float64) float64 {
// 	return (b * b) - (4 * a * c)
// }
