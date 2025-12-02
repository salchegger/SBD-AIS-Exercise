package client

import (
	"context"
	pb2 "exc8/exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb2.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.Dial(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb2.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()

	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	fmt.Println("Available drinks:")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:\"%s\"  price:%d  description:\"%s\"\n", d.Id, d.Name, d.Price, d.Description)
	}

	// 2. Order first round
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	firstRound := map[int32]int32{
		1: 2,
		2: 2,
		3: 2,
	}
	for id, qty := range firstRound {
		item := &pb2.OrderItem{DrinkId: id, Quantity: qty}
		if _, err := c.client.OrderDrink(ctx, &pb2.OrderRequest{Item: item}); err != nil {
			return err
		}
		fmt.Printf("\t> Ordering: %d x %s\n", qty, getDrinkName(id, drinksResp.Drinks))
	}

	// 3. Order second round
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	secondRound := map[int32]int32{
		1: 6,
		2: 6,
		3: 6,
	}
	for id, qty := range secondRound {
		item := &pb2.OrderItem{DrinkId: id, Quantity: qty}
		if _, err := c.client.OrderDrink(ctx, &pb2.OrderRequest{Item: item}); err != nil {
			return err
		}
		fmt.Printf("\t> Ordering: %d x %s\n", qty, getDrinkName(id, drinksResp.Drinks))
	}

	// 4. Get total orders
	fmt.Println("Getting the bill 💹💹💹")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	// Aggregate totals
	totals := make(map[int32]int32)
	for _, o := range ordersResp.Orders {
		totals[o.DrinkId] += o.Quantity
	}

	for id, qty := range totals {
		fmt.Printf("\t> Total: %d x %s\n", qty, getDrinkName(id, drinksResp.Drinks))
	}

	return nil
}

// helper to get drink name from ID
func getDrinkName(id int32, drinks []*pb2.Drink) string {
	for _, d := range drinks {
		if d.Id == id {
			return d.Name
		}
	}
	return "Unknown"
}
