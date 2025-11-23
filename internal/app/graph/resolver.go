package graph

import (
	contentv1 "gqlgen-playground/internal/pb/content/v1"

	"google.golang.org/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	contentClient contentv1.ContentServiceClient
}

func NewResolver(conn *grpc.ClientConn) *Resolver {
	return &Resolver{
		contentClient: contentv1.NewContentServiceClient(conn),
	}
}
