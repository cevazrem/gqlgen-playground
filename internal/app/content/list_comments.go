package content

import (
	"context"
	"errors"
	"fmt"
	desc "gqlgen-playground/internal/pb/content/v1"
	content_dto "gqlgen-playground/internal/pkg/dto/content"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListComments получить список комментариев
func (i *Implementation) ListComments(ctx context.Context, req *desc.ListCommentsRequest) (*desc.ListCommentsResponse, error) {
	if err := validateListCommentsRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateListCommentsRequest: %v", err)
	}

	comments, total, err := i.contentService.ListComments(ctx, content_dto.BuildCommentFilterByListReq(req.Filter), content_dto.BuildLimitOffsetPagination(req.Page, req.PerPage))
	if err != nil {
		return nil, fmt.Errorf("i.contentService.ListComments: %w", err)
	}

	protoComments := make([]*desc.Comment, 0, len(comments))
	for _, comment := range comments {
		protoComments = append(protoComments, comment.ToProto())
	}

	return &desc.ListCommentsResponse{
		Comments: protoComments,
		Total:    total,
	}, nil
}

func validateListCommentsRequest(req *desc.ListCommentsRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return nil
}
