package content

import (
	"context"
	"errors"
	"fmt"
	desc "gqlgen-playground/internal/pb/content/v1"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser получить данные пользователя
func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	if err := validateGetUserRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateGetUserRequest: %v", err)
	}

	user, err := i.contentService.GetUser(ctx, content_dto.BuildUserFilterByGetReq(req))
	if err != nil {
		return nil, fmt.Errorf("i.contentService.GetUser: %w", err)
	}

	return &desc.GetUserResponse{User: user.ToProto()}, nil
}

func validateGetUserRequest(req *desc.GetUserRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	if req.Filter == nil {
		return errors.New("filter is required")
	}

	switch f := req.Filter.(type) {
	case *desc.GetUserRequest_Id:
		id := strings.TrimSpace(f.Id)
		if id == "" {
			return errors.New("id must not be empty")
		}
		if _, err := uuid.Parse(id); err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}

	case *desc.GetUserRequest_Username:
		username := strings.TrimSpace(f.Username)
		if username == "" {
			return errors.New("username must not be empty")
		}

	default:
		return errors.New("unknown filter type")
	}

	return nil
}
