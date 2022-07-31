package config

import (
	"context"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	aClient "github.com/voltgizerz/public-grpc/auth/gen"
	oClient "github.com/voltgizerz/public-grpc/order/gen"
	uClient "github.com/voltgizerz/public-grpc/user/gen"
)

// GRPCClient - hold data client grpc.
type GRPCClient struct {
	OrderClient oClient.OrderServiceClient
	UserClient  uClient.UserServiceClient
	AuthClient  aClient.AuthServiceClient
}

// InitGRPC - return client gRPC.
func InitGRPC(ctx context.Context) *GRPCClient {
	log := SetupLog()

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

	authClient, err := grpc.DialContext(ctx, os.Getenv("AUTH_GRPC"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect to auth gRPC: %v", err)
	}
	// defer userClient.Close()

	return &GRPCClient{
		OrderClient: oClient.NewOrderServiceClient(orderClient),
		UserClient:  uClient.NewUserServiceClient(userClient),
		AuthClient:  aClient.NewAuthServiceClient(authClient),
	}
}
