package content

import (
	"context"
	"errors"
	"fmt"
	desc "gqlgen-playground/internal/pb/content/v1"
	content_dto "gqlgen-playground/internal/pkg/dto/content"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPost Получить информацию о посте
func (i *Implementation) GetPost(ctx context.Context, req *desc.GetPostRequest) (*desc.GetPostResponse, error) {
	if err := validateGetPostRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateGetPostRequest: %v", err)
	}

	post, err := i.contentService.GetPost(ctx, content_dto.PostFilter{IdsIn: []string{req.Id}})
	if err != nil {
		return nil, fmt.Errorf("i.contentService.GetPost: %w", err)
	}

	return &desc.GetPostResponse{Post: post.ToProto()}, nil
}

func validateGetPostRequest(req *desc.GetPostRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required, is.UUID),
	)
}
