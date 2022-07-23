package client

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	oClient "github.com/voltgizerz/public-grpc/order/gen"
	uClient "github.com/voltgizerz/public-grpc/user/gen"
)

// GrpcClient - hold data client grpc.
type GrpcClient struct {
	OrderClient oClient.OrderServiceClient
	UserClient  uClient.UserServiceClient
}

// InitGRPC - return client gRPC.
func InitGRPC(ctx context.Context) *GrpcClient {
	orderClient, err := grpc.DialContext(ctx, os.Getenv("ORDER_GRPC"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect to order gRPC: %v", err)
	}
	// defer orderClient.Close()

	userClient, err := grpc.DialContext(ctx, os.Getenv("USER_GRPC"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect to user gRPC: %v", err)
	}
	// defer userClient.Close()

	return &GrpcClient{
		OrderClient: oClient.NewOrderServiceClient(orderClient),
		UserClient:  uClient.NewUserServiceClient(userClient),
	}
}
