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

// CreateComment создает новый комментарий
func (i *Implementation) CreateComment(ctx context.Context, req *desc.CreateCommentRequest) (*desc.CreateCommentResponse, error) {
	if err := validateCreateCommentRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateCreateCommentRequest: %v", err)
	}

	comment, err := i.contentService.CreateComment(ctx, content_model.Comment{
		PostID:   req.PostId,
		AuthorID: req.AuthorId,
		Body:     req.Body,
	})
	if err != nil {
		return nil, fmt.Errorf("i.contentService.CreateComment: %w", err)
	}

	return &desc.CreateCommentResponse{Comment: comment.ToProto()}, nil
}

func validateCreateCommentRequest(req *desc.CreateCommentRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return validation.ValidateStruct(req,
		validation.Field(&req.PostId, validation.Required, is.UUID),
		validation.Field(&req.AuthorId, validation.Required, is.UUID),
		validation.Field(&req.Body, validation.Required),
	)
}
