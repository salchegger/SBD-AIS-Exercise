package server

import (
	"context"
	pb2 "exc8/exc8/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb2.UnimplementedOrderServiceServer
}

// Prepopulate drinks and in-memory orders
var drinks = []*pb2.Drink{
	{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
	{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
	{Id: 3, Name: "Coffee", Price: 1, Description: "Mifare isn't that secure"},
}

var orders = []*pb2.OrderItem{}

func StartGrpcServer() error {
	srv := grpc.NewServer()
	grpcService := &GRPCService{}
	pb2.RegisterOrderServiceServer(srv, grpcService)

	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}

	fmt.Println("gRPC server running on :4000")
	return srv.Serve(lis)
}

// GetDrinks returns all preloaded drinks
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb2.DrinkList, error) {
	return &pb2.DrinkList{Drinks: drinks}, nil
}

// OrderDrink stores an order in memory
func (s *GRPCService) OrderDrink(ctx context.Context, req *pb2.OrderRequest) (*wrapperspb.BoolValue, error) {
	orders = append(orders, req.Item)
	return wrapperspb.Bool(true), nil
}

// GetOrders returns all orders made so far
func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb2.AllOrders, error) {
	return &pb2.AllOrders{Orders: orders}, nil
}

// helper function to get drink name by ID
func getDrinkName(id int32) string {
	for _, d := range drinks {
		if d.Id == id {
			return d.Name
		}
	}
	return "Unknown Drink"
}
