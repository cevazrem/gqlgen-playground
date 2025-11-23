package content

import (
	"context"
	"errors"
	"fmt"
	desc "gqlgen-playground/internal/pb/content/v1"
	content_model "gqlgen-playground/internal/pkg/model/content"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePost создать новый пост
func (i *Implementation) CreatePost(ctx context.Context, req *desc.CreatePostRequest) (*desc.CreatePostResponse, error) {
	if err := validateCreatePostRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateCreatePostRequest: %v", err)
	}

	post, err := i.contentService.CreatePost(ctx, content_model.Post{
		AuthorID: req.AuthorId,
		Title:    req.Title,
		Body:     req.Body,
	})
	if err != nil {
		return nil, fmt.Errorf("i.contentService.CreatePost: %w", err)
	}

	return &desc.CreatePostResponse{Post: post.ToProto()}, nil
}

func validateCreatePostRequest(req *desc.CreatePostRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return validation.ValidateStruct(req,
		validation.Field(&req.AuthorId, validation.Required, is.UUID),
		validation.Field(&req.Title, validation.Required),
		validation.Field(&req.Body, validation.Required),
	)
}
