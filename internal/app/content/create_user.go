package content

import (
	"context"
	"errors"
	"fmt"
	desc "gqlgen-playground/internal/pb/content/v1"
	content_model "gqlgen-playground/internal/pkg/model/content"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/guregu/null.v4/zero"
)

// CreateUser создает нового пользователя
func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if err := validateCreateUserRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateCreateUserRequest: %v", err)
	}

	user, err := i.contentService.CreateUser(ctx, content_model.User{
		Name:     req.GetName(),
		Username: req.GetUsername(),
		Phone:    zero.StringFromPtr(req.Phone),
		Email:    zero.StringFromPtr(req.Email),
	})
	if err != nil {
		return nil, fmt.Errorf("i.contentService.CreateUser: %w", err)
	}

	return &desc.CreateUserResponse{User: user.ToProto()}, nil
}

func validateCreateUserRequest(req *desc.CreateUserRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return validation.ValidateStruct(req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Username, validation.Required),
	)
}
