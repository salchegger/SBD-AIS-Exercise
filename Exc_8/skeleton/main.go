package main

import (
	"exc8/client"
	"exc8/server"
	"log"
	"time"
)

func main() {
	// Start server in goroutine
	go func() {
		if err := server.StartGrpcServer(); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Give server time to start
	time.Sleep(1 * time.Second)

	// Start client
	grpcClient, err := client.NewGrpcClient()
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	if err := grpcClient.Run(); err != nil {
		log.Fatalf("Client run failed: %v", err)
	}

	println("Orders complete!")
}
