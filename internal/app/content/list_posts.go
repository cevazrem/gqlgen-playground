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

// ListPosts получить список постов
func (i *Implementation) ListPosts(ctx context.Context, req *desc.ListPostsRequest) (*desc.ListPostsResponse, error) {
	if err := validateListPostsRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateListPostsRequest: %v", err)
	}

	posts, total, err := i.contentService.ListPosts(ctx, content_dto.BuildPostFilterByListReq(req.Filter), content_dto.BuildLimitOffsetPagination(req.Page, req.PerPage))
	if err != nil {
		return nil, fmt.Errorf("i.contentService.ListPosts: %w", err)
	}

	protoPosts := make([]*desc.Post, 0, len(posts))
	for _, post := range posts {
		protoPosts = append(protoPosts, post.ToProto())
	}

	return &desc.ListPostsResponse{
		Posts: protoPosts,
		Total: total,
	}, nil
}

func validateListPostsRequest(req *desc.ListPostsRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return nil
}
