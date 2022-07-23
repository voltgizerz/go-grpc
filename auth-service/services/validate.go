package services

import (
	"context"
	"errors"
	"log"

	pb "github.com/voltgizerz/public-grpc/auth/gen"
	"google.golang.org/grpc/metadata"
)

// Validate - validate user.
func (s *Server) Validate(ctx context.Context, in *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	log.Printf("Received: %v", in)
	
	token := ""
	if in.Token != ""	{
		token = in.Token
	}

	if token == "" {
		tokenMetaData, err := GetAuthorizationMetaData(ctx)
		if err != nil {
			return &pb.ValidateResponse{
				Status: 401,
				Error: err.Error(),
			}, nil
		}
		token = tokenMetaData
	}

	claims, err:= s.Jwt.ValidateToken(token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: 401,
			Error: err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: 200,
		UserId: claims.ID,
	}, nil
}

// GetAuthorizationMetaData - get authorization meta data.
func GetAuthorizationMetaData(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
  token := md.Get("authorization")[0]

	if !ok {
		return "", errors.New("missing metadata")
	}
	return token, nil
}
